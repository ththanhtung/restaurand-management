package routes

import (
	"mongotest/controllers"
	"mongotest/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/users", middlewares.RequireAuth() ,controllers.GetUsers())
	incommingRoute.GET("/users/:id", controllers.GetUser())
	incommingRoute.POST("/users", controllers.NewUser())
	incommingRoute.PATCH("/users/:id", controllers.UpdateUser())
	incommingRoute.DELETE("/users/:id", controllers.DeleteUser())
}
