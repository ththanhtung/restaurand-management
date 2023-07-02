package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/orders")
	incommingRoute.GET("/orders/:id")
	incommingRoute.POST("/orders", controllers.NewOrder())
	incommingRoute.PATCH("/orders/:id")
	incommingRoute.DELETE("/orders/:id")
}
