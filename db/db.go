package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionTimeout = 10
	uri               = "mongodb://testUser:testPassword@0.0.0.0:27017"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
	}

	return client, ctx, cancel
}
