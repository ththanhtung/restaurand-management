package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/orderItem")
	incommingRoute.GET("/orderItem/:id")
	incommingRoute.POST("/orderitems", controllers.NewOrderItem())
	incommingRoute.PATCH("/orderitems/:id", controllers.UpdateOrderItem())
	incommingRoute.DELETE("/orderItem/:id")
}
