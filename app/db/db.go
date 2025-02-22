package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var URI = os.Getenv("MONGODB_URI")

func GetClient() (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	return client, nil
}