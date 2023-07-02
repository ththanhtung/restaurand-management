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

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "tables")

func NewTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tableReq *models.TableRequest

		err := c.ShouldBindJSON(&tableReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while binding form",
			})
			return
		}

		if err := tableReq.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		tableCount, err := tableCollection.CountDocuments(ctx, bson.M{"number": tableReq.Number})
		if tableCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "table already exist",
			})
			return
		}
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while checking table",
			})
			return
		}

		table := &models.Table{
			Number: tableReq.Number,
		}
		table.ID = primitive.NewObjectID()
		table.TableId = table.ID.Hex()

		// save data to mongoDB
		_, err = tableCollection.InsertOne(ctx, table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while inserting new table",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, table)
	}
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cursor, err := tableCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error retrieving tables",
			})
			return
		}
		defer cancel()

		var results []models.Table

		if err := cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error binding tables",
			})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}