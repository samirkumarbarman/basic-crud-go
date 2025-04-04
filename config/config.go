package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://samirkumar:220nipu@db.dho4l.mongodb.net/go_mongodb_crud?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify that the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("go_mongodb_crud").Collection(collectionName)
}
