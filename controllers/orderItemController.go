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

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItems")

func NewOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItemReq *models.OrderItemRequest
		c.ShouldBindJSON(&orderItemReq)

		if err := orderItemReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		orderItem := &models.OrderItem{
			OrderId:  orderItemReq.OrderId,
			FoodId:   orderItemReq.FoodId,
			Quantity: orderItemReq.Quantity,
		}
		orderItem.ID = primitive.NewObjectID()
		orderItem.OrderItemId = orderItem.ID.Hex()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := orderItemCollection.InsertOne(ctx, orderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while inserting order item",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, orderItem)
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemIdReq := c.Param("id")
		orderItemId, _ := primitive.ObjectIDFromHex(orderItemIdReq)

		var orderItemReq *models.OrderItemUpdateRequest

		c.ShouldBindJSON(&orderItemReq)

		if err := orderItemReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var updatedObject primitive.D = primitive.D{}

		if orderItemReq.FoodId != "" {
			updatedObject = append(updatedObject, bson.E{"foodid", orderItemReq.FoodId})
		}
		if orderItemReq.OrderId != "" {
			updatedObject = append(updatedObject, bson.E{"orderid", orderItemReq.OrderId})
		}
		if orderItemReq.Quantity != 0 {
			updatedObject = append(updatedObject, bson.E{"quantity", orderItemReq.Quantity})
		}

		var updatedOrderItem *models.OrderItem
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := orderItemCollection.FindOneAndUpdate(ctx, bson.M{"_id": orderItemId}, bson.D{{"$set", updatedObject}}).Decode(&updatedOrderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, updatedOrderItem)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemIdReq := c.Param("id")
		orderItemId, _ := primitive.ObjectIDFromHex(orderItemIdReq)
		var orderItem *models.OrderItem

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := orderItemCollection.FindOne(ctx, bson.M{"_id": orderItemId}).Decode(&orderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while fetching order item",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, orderItem)
	}
}
