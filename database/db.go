package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string = os.Getenv("MONGO_URI")

func DBInstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

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