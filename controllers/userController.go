package controllers

import (
	"context"
	"log"
	"mongotest/database"
	"mongotest/helpers"
	"mongotest/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")

		c.JSON(http.StatusOK, user)
	}
}

func NewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		c.ShouldBindJSON(&user)
		
		// validate user stuct to make sure user input is correct
		err := user.Validate(); 
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// check if email already exist
		userCount, err := userCollection.CountDocuments(ctx, bson.D{{"email", user.Email}})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error when checking email",
			})
			return
		}

		if userCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "email already exist",
			})
			return
		}

		// hash password
		password := helpers.HashPassword(*user.Password)
		user.Password = &password

		// add created at, updated at, id for user
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		// Generate token for user
		token, _ := helpers.GenerateToken(*user.Email, user.UserID)
		user.Token = &token

		// insert user to database
		_, err = userCollection.InsertOne(ctx, user)
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"user was not created",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context) {
		var user *models.User
		var foundUser *models.User

		c.ShouldBindJSON(&user)

		if err := user.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		err := userCollection.FindOne(ctx, bson.D{{"email", user.Email}}).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occurred while reading user",
			})
			return
		}

		passwordMatch, errMsg := helpers.VerifyPassword(*foundUser.Password, *user.Password)
		if passwordMatch != true {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errMsg,
			})
			return
		}

		token,_ := helpers.GenerateToken(*foundUser.Email, *&foundUser.UserID)

		helpers.UpdateToken(foundUser.UserID, token)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
