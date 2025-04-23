package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var URI = os.Getenv("MONGODB_URI")

func GetClient() (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(URI))
}

func DisconnectClient(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}