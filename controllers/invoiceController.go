package controllers

// import (
// 	"mongotest/models"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func NewInvoice() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var invoiceReq *models.InvoiceRequest

// 		c.ShouldBindJSON(&invoiceReq)

// 		invoiceLookupStage := bson.D{{"$lookup",bson.D{{"from","orderitems"}, {"localField","orderid"}, {"foreignField","orderid"},{"as", "orderitem"}}}}
// 		invoiceUnwindStage := bson.D{{"$unwind","$orderitem"}}
// 		invoiceProjectStage := bson.D{{"$project", bson.D{{"invoiceid",1}, {"foodid", "$orderitem.foodid"},{"quantity","$orderitem.quantity"}}}}
// 		orderItemLookupStage := bson.D{{"$lookup", bson.D{{"from","foods"},{"localField","foodid"},{"foreignField","foodid"},{"as","food"}}}}
// 		orderItemUnwindStage := bson.D{{"$unwind","$food"}}
// 		orderItemAddFieldStage := bson.D{{"$addField", bson.D{""} bson.D{{"$multiply", []interface{}{
// 			"$food.price", "$quantity",
// 		}}}}}
// 		orderItemProjectStage := bson.D{{"$project", bson.D{{"invoiceid",1},}}}
// 	}
// }