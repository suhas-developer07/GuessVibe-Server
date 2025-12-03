package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning unable to find .env file")
	}
	MongoDb := os.Getenv("MONGODB_URI")
	if MongoDb == "" {
		log.Fatal("MONGODB_URI not found in environment variables")
	}
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	clientoptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.NewClient(clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client

}
func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME: ", databaseName)
	collection := client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}
	return collection
}
