package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectMongoDB(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		if err != nil {
			log.Fatalf("MongoDB connection error: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		if err != nil {
			log.Fatalf("Error during connection MongoDB: %v", err)
		}
	}

	fmt.Println("MongoDB Connected")
	DB = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("list_service_db").Collection(collectionName)
}
