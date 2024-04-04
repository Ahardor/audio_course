package test

import (
	"context"
	"iotvisual/processor/internal/queries"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTestCollection() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAuth(options.Credential{
				Username: "iotvisual",
				Password: "iotvisualpass",
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("test").Collection("testcollection")
}

func NewQueries(c *mongo.Collection) *queries.Queries {
	return queries.New(c)
}
