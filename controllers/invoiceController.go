package controllers

import (
	"context"
	"mongotest/database"
	"mongotest/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
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