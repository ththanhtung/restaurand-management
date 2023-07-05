package controllers

import (
	"context"
	"io/ioutil"

	"log"
	"mongotest/database"
	"mongotest/models"
	"net/http"
	"regexp"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "foods")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cursor, err := foodCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error retrieving foods",
			})
			return
		}
		defer cancel()

		var results []models.Food

		if err := cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error binding foods",
			})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId := c.Param("id")

		var food *models.Food

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := foodCollection.FindOne(ctx, bson.D{{"foodid", foodId}}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while checking food",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, food)
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

		// Match .jpg or .png

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

		// check if the file extension is valid
		validExt := regexp.MustCompile(`\.(jpg|png)$`)
		if !validExt.MatchString(file.Filename) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "only allow to upload jpg and png file",
			})
			return
		}

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
			Name:  foodReq.Name,
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
	return func(c *gin.Context) {
		foodIdReq := c.Param("id")
		foodId, _ := primitive.ObjectIDFromHex(foodIdReq)

		var foodReq *models.FoodRequest
		var updatedFood *models.Food

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
		err = foodCollection.FindOne(ctx, bson.D{{"_id", foodId}}).Decode(&updatedFood)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error checking food",
			})
			return
		}
		defer cancel()

		updatedFood.Name = foodReq.Name
		updatedFood.Price = foodReq.Price
		if foodReq.Image != nil && foodReq.Image.Filename != "" {
			log.Println("saving image file")
		}

		// update document to db
		upsert := true
		options := options.UpdateOptions{
			Upsert: &upsert,
		}
		_, err = foodCollection.UpdateOne(ctx, bson.M{"_id": foodId}, bson.M{"$set": updatedFood}, &options)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while updating food:" + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, updatedFood)
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodIdReq := c.Param("id")
		foodId, _ := primitive.ObjectIDFromHex(foodIdReq)
		var food *models.Food

		filter := bson.M{"_id": foodId}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := foodCollection.FindOneAndDelete(ctx, filter).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, food)
	}
}

func GetFoodImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId := c.Param("id")

		var food *models.Food

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := foodCollection.FindOne(ctx, bson.D{{"foodid", foodId}}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while checking food",
			})
			return
		}
		defer cancel()

		c.Data(http.StatusOK, "image/jpg", food.Image)
	}
}
