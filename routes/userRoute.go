package routes

import (
	"mongotest/controllers"
	"mongotest/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoute *gin.Engine) {
	userRoutes := incommingRoute.Group("/users", middlewares.RequireAuth())
	userRoutes.GET("/", controllers.GetUser())
	userRoutes.PATCH("/", controllers.UpdateUser())
	userRoutes.DELETE("/", controllers.DeleteUser())
}
