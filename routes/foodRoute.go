package routes

import (
	"mongotest/controllers"
	"mongotest/middlewares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoute *gin.Engine){
	incommingRoute.GET("/foods", controllers.GetFoods())
	incommingRoute.GET("/foods/:id", controllers.GetFood())
	incommingRoute.POST("/foods", controllers.NewFood())
	incommingRoute.PATCH("/foods/:id",middlewares.RequireAuth(), controllers.UpdateFood())
	incommingRoute.DELETE("/foods/:id",middlewares.RequireAuth(), controllers.DeleteFood())
}