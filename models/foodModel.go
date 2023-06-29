package models

import (
	"mime/multipart"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID     primitive.ObjectID `bson:"_id"`
	FoodId string
	Name   *string  `form:"name"`
	Price  *float32 `form:"price"`
	Image  []byte
}

type FoodRequest struct {
	Name  *string               `form:"name" validate:"required,min=1"`
	Price *float32              `form:"price" validate:"required"`
	Image *multipart.FileHeader `form:"foodImage"`
}

func (f *FoodRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
