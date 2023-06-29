package middlewares

import (
	"context"
	"mongotest/database"
	"mongotest/helpers"
	"mongotest/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func RequireAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"no authenticate token provided",
			})
			c.Abort()
			return
		}

		claims, errMsg := helpers.VerifyToken(token)

		if errMsg != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":errMsg})
			c.Abort()
			return
		}

		var user *models.User

		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

		err := userCollection.FindOne(ctx, bson.D{{"email", claims.Email}}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"error occurred while reading user",
			})
			return
		}

		c.Set("user", user)
	
		c.Next()
	}
}