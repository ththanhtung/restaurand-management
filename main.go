package main

import (
	"mongotest/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":"+ port)
}