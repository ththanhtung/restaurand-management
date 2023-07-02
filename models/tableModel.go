package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID      primitive.ObjectID `bson:"_id"`
	TableId string
	Number  int
}

type TableRequest struct {
	Number int `json:"tableNumber" validate:"required"`
}

func (t *TableRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
