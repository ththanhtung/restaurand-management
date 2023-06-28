package routes

import "github.com/gin-gonic/gin"

func TableRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/tables")
	incommingRoute.GET("/tables/:id")
	incommingRoute.POST("/tables")
	incommingRoute.PATCH("/tables/:id")
	incommingRoute.DELETE("/tables/:id")
}
