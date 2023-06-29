package controllers

import (
	"context"
	"io/ioutil"
	"mongotest/database"
	"mongotest/models"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "foods")

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func NewFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foodReq *models.FoodRequest

		// single file
		err := c.ShouldBind(&foodReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while binding form",
			})
			return
		}

		if err := foodReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		foodCount, err := foodCollection.CountDocuments(ctx, bson.D{{"name", foodReq.Name}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while checking food name",
			})
			return
		}
		defer cancel()

		if foodCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "food already exist",
			})
			return
		}

		// get the file from the form
		file := foodReq.Image

		// open file to take data
		data, err := file.Open()
		defer data.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while opening image",
			})
			return
		}

		// get binary file from the data to save to mongoDB
		fileBytes, err := ioutil.ReadAll(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while reading image",
			})
			return
		}

		// 
		food := &models.Food{
			Name: foodReq.Name,
			Price: foodReq.Price,
			Image: fileBytes,
		}
		food.ID = primitive.NewObjectID()
		food.FoodId = food.ID.Hex()

		// save data to mongoDB
		_, err = foodCollection.InsertOne(ctx, food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while inserting new food",
			})
			return
		}

		c.JSON(http.StatusCreated, food)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func DeleteFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
