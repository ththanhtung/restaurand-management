package routes

import "github.com/gin-gonic/gin"

func OrderItemRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/orderItem")
	incommingRoute.GET("/orderItem/:id")
	incommingRoute.POST("/orderItem")
	incommingRoute.PATCH("/orderItem/:id")
	incommingRoute.DELETE("/orderItem/:id")
}
