package database

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoDB(uri string) *Mongo {
	// Create a new client and connect to the mongodb server
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to mongoDB:", err)
	}

	// Create a context with timeout for connecting to mongodb
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal("Failed to ping mongoDB:", err)
	// }
	return &Mongo{Client: client}
}
