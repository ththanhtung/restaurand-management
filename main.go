package main

import (
	"log"
	"mongotest/initializers"
	"mongotest/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("init")
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	if port == "" {
		port = "8080"
	}

	router.Use(cors.New(config))

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)
	routes.UserRoutes(router)
	routes.TableRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":" + port)
}
