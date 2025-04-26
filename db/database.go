package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	dbPass := os.Getenv("DBPASS")
	if dbPass == "" {
		log.Fatal("DBPASS environment variable is not set")
	}

	uri := fmt.Sprintf("mongodb+srv://analytics:%s@amlak.wjtlb.mongodb.net/amlak?retryWrites=true&w=majority", dbPass)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected and pinged.")
	return client

}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Amlak").Collection(collectionName)

	return collection
}
