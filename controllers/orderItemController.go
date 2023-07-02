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

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItems")

func NewOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var orderItemReq *models.OrderItemRequest
		c.ShouldBindJSON(&orderItemReq)

		if err := orderItemReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		orderItem := &models.OrderItem{
			OrderId: orderItemReq.OrderId,
			FoodId: orderItemReq.FoodId,
			Quantity: orderItemReq.Quantity,
		}
		orderItem.ID = primitive.NewObjectID()
		orderItem.OrderItemId = orderItem.ID.Hex()
		
		ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
		_, err := orderItemCollection.InsertOne(ctx, orderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"error occurred while inserting order item",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, orderItem)
	}
}