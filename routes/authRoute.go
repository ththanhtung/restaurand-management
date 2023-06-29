package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoute *gin.Engine) {
	incommingRoute.POST("/auth/signup", controllers.NewUser())
	incommingRoute.POST("/auth/login", controllers.Login())
}