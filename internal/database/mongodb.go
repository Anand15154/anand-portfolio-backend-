package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(uri, databaseName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Attempting to connect to MongoDB...")
	log.Printf("Database name: %s", databaseName)

	// Log URI status without exposing credentials
	if uri == "mongodb://localhost:27017" {
		log.Printf("Using localhost MongoDB URI")
	} else {
		log.Printf("Using custom MongoDB URI (Atlas/remote)")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Failed to create MongoDB client: %v", err)
		return nil, err
	}

	// Ping the database with timeout
	log.Printf("Pinging MongoDB database...")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		log.Printf("Please check your MONGODB_URI environment variable")
		return nil, err
	}

	log.Println("✅ Connected to MongoDB successfully")

	return &MongoDB{
		Client:   client,
		Database: client.Database(databaseName),
	}, nil
}

func (m *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}

func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return m.Database.Collection(collectionName)
}
