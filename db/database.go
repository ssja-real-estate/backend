package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	rawPass := os.Getenv("DBPASS")
	fmt.Print("get pass in env")
	if rawPass == "" {
		log.Fatal("DBPASS environment variable is not set")
	}
	pass := url.QueryEscape(rawPass)

	uri := "mongodb+srv://analytics:" + pass +
		"@amlak.wjtlb.mongodb.net/amlak?retryWrites=true&w=majority&appName=Amlak"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Successfully connected and pinged.")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("amlak").Collection(collectionName)
}
