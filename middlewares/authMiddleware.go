package middlewares

import (
	"mongotest/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == ""{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"no authenticate token provided",
			})
			c.Abort()
			return
		}

		claims, err := helpers.VerifyToken(token)

		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}