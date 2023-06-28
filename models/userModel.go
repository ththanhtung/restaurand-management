package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     string             `json:"userId"`
	Username   *string            `json:"username" validate:"required,min=2"`
	Password   *string            `json:"password" validate:"required,min=8"`
	Email      *string            `json:"email" validate:"email,required"`
	Phone      *int               `json:"phone"`
	Address    *string            `json:"address"`
	Token      *string            `json:"token"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
