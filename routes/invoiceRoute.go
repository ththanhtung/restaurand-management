package routes

import (
	"mongotest/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingRoute *gin.Engine) {
	incommingRoute.GET("/invoices")
	incommingRoute.GET("/invoices/:id")
	incommingRoute.POST("/invoices", controllers.NewInvoice())
	incommingRoute.PATCH("/invoices/:id")
	incommingRoute.DELETE("/invoices/:id")
}