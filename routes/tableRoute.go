package routes

import (
	"mongotest/controllers"
	"mongotest/middlewares"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/tables", controllers.GetTables())
	incommingRoute.GET("/tables/:id")
	incommingRoute.POST("/tables", middlewares.RequireAuth(), controllers.NewTable())
	incommingRoute.PATCH("/tables/:id")
	incommingRoute.DELETE("/tables/:id")
}
