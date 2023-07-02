package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID      primitive.ObjectID `bson:"_id"`
	OrderId string
	UserId  string
	TableId string
}

type OrderRequest struct {
	OrderId string `json:"orderId"`
	UserId  string `json:"userId"`
	TableId string `json:"tableId"`
}

func (o *OrderRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
