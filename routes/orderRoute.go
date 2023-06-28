package routes

import "github.com/gin-gonic/gin"

func OrderRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/orders")
	incommingRoute.GET("/orders/:id")
	incommingRoute.POST("/orders")
	incommingRoute.PATCH("/orders/:id")
	incommingRoute.DELETE("/orders/:id")
}
