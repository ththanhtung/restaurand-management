package main

import (
	"log"
	"mongotest/initializers"
	"mongotest/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init(){
	log.Println("init")
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)
	routes.UserRoutes(router)
	routes.TableRoutes(router)

	router.Run(":"+ port)
}