package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/orderItem")
	incommingRoute.GET("/orderItem/:id")
	incommingRoute.POST("/orderitems", controllers.NewOrderItem())
	incommingRoute.PATCH("/orderItem/:id")
	incommingRoute.DELETE("/orderItem/:id")
}
