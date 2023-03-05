package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error

type Collection string

const (
	ProductsCollection Collection = "products"
)

// * reference the value
var mongoOnce sync.Once

const (
	URI      = "mongodb://localhost:27017"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(URI)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}

// & point the block of memory that store this object
