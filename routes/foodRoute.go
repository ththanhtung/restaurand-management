package routes

import (
	"mongotest/controllers"
	"mongotest/middlewares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoute *gin.Engine){
	// limit file size to less than 8MB
	incommingRoute.MaxMultipartMemory = 8 << 20
	
	incommingRoute.GET("/foods", controllers.GetFoods())
	incommingRoute.GET("/foods/:id", controllers.GetFood())
	incommingRoute.GET("/foods/:id/image", controllers.GetFoodImage())
	incommingRoute.POST("/foods", middlewares.RequireAuth(), controllers.NewFood())
	incommingRoute.PATCH("/foods/:id",middlewares.RequireAuth(), controllers.UpdateFood())
	incommingRoute.DELETE("/foods/:id",middlewares.RequireAuth(), controllers.DeleteFood())
}