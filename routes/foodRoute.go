package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoute *gin.Engine){
	incommingRoute.GET("/foods", controllers.GetFoods())
	incommingRoute.GET("/foods/:id", controllers.GetFood())
	incommingRoute.POST("/foods", controllers.NewFood())
	incommingRoute.PATCH("/foods/:id", controllers.UpdateFood())
	incommingRoute.DELETE("/foods/:id", controllers.DeleteFood())
}