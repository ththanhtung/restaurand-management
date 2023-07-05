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
	Quantity    int
	UnitPrice   float32
}

type OrderItemRequest struct {
	OrderId  string `json:"orderId" validate:"required"`
	FoodId   string `json:"foodId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type OrderItemUpdateRequest struct {
	OrderId  string `json:"orderId"`
	FoodId   string `json:"foodId"`
	Quantity int    `json:"quantity"`
}

func (o OrderItemRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

func (o OrderItemUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
