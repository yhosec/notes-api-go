package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo() (*mongo.Database) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return client.Database("node")
}
