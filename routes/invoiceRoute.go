package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/invoices")
	incommingRoute.GET("/invoices/:id")
	incommingRoute.POST("/invoices")
	incommingRoute.PATCH("/invoices/:id")
	incommingRoute.DELETE("/invoices/:id")
}