package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() (*mongo.Client, error) {
	// MongoDB Atlas connection string
uri :=os.Getenv("MONGO_URI")
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB Atlas
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB Atlas!")
	return client, nil
}