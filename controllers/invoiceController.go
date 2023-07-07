package controllers

import (
	"context"
	"mongotest/database"
	"mongotest/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var invoicesCollection *mongo.Collection = database.OpenCollection(database.Client, "invoices")

func NewInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoiceReq *models.InvoiceRequest

		c.ShouldBindJSON(&invoiceReq)

		invoice := &models.Invoice{
			OrderId: invoiceReq.OrderId,
		}
		invoice.ID = primitive.NewObjectID()
		invoice.InvoiceId = invoice.ID.Hex()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := invoicesCollection.InsertOne(ctx, invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, invoice)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceIdReq := c.Param("id")
		invoiceId, err := primitive.ObjectIDFromHex(invoiceIdReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var invoice *models.Invoice
		invoicesCollection.FindOne(ctx, bson.M{"_id": invoiceId}).Decode(&invoice)
		defer cancel()

		invoicePipeline := bson.A{
			bson.D{{"$match", bson.D{{"orderid", invoice.OrderId}}}},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "foods"},
						{"localField", "foodid"},
						{"foreignField", "foodid"},
						{"as", "food"},
					},
				},
			},
			bson.D{{"$unwind", bson.D{{"path", "$food"}}}},
			bson.D{
				{"$project",
					bson.D{
						{"foodname", "$food.name"},
						{"quantity", "$quantity"},
						{"orderid", "$orderid"},
						{"unitprice", "$unitprice"},
					},
				},
			},
			bson.D{
				{"$group",
					bson.D{
						{"_id", "$orderid"},
						{"foods",
							bson.D{
								{"$push",
									bson.D{
										{"foodname", "$foodname"},
										{"quantity", "$quantity"},
										{"unitPrice", "$unitprice"},
									},
								},
							},
						},
						{"total", bson.D{{"$sum", "$unitprice"}}},
					},
				},
			},
		}

		cursor, err := orderItemCollection.Aggregate(ctx, invoicePipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		defer cancel()

		results := []primitive.M{}
		if err = cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, results[0])
	}
}