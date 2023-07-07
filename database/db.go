package database

import (
	"context"
	"log"
	"mongotest/initializers"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	initializers.LoadEnv()

	var mongoURI string = os.Getenv("MONGO_URI")
	log.Println("mongo uri",mongoURI)
	
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err.Error())
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err.Error())
	}

	defer cancel()

	log.Println("connected to mongo")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	databaseName := os.Getenv("DATABASE_NAME")
	collection := client.Database(databaseName).Collection(collectionName)

	return collection
}