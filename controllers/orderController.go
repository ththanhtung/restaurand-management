package controllers

import (
	"context"
	"mongotest/database"
	"mongotest/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func NewOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var orderReq *models.OrderRequest

		c.ShouldBindJSON(&orderReq)

		order := &models.Order{
			UserId: orderReq.UserId,
			TableId: orderReq.TableId,
		}
		order.ID = primitive.NewObjectID()
		order.OrderId = order.ID.Hex()

		ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
		_, err := orderCollection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return 			
		}
		defer cancel()
		 
		c.JSON(http.StatusCreated, order)
	}
}

