package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invoice struct {
	ID        primitive.ObjectID `bson:"_id"`
	InvoiceId string
	OrderId   string
	Total     string
}

type InvoiceRequest struct {
	OrderId string `json:"orderId"`
}