package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string = "mongodb://localhost:27017"

func DBInstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("mongo uri",mongoURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err.Error())
	}

	defer cancel()

	log.Println("connected to mongo")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	collection := client.Database("restaurantDB").Collection(collectionName)

	return collection
}