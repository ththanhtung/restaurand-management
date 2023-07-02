package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrderItemId string
	OrderId     string
	FoodId      string
	Quantity    string
}

type OrderItemRequest struct {
	OrderId  string `json:"orderId" validate:"required"`
	FoodId   string `json:"foodId" validate:"required"`
	Quantity string `json:"quantity" validate:"required"`
}

func (o OrderItemRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
