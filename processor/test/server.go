package test

import (
	"context"
	"iotvisual/processor/internal/pkg/queries"
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
			SetAuth(options.Credential{
				Username: "iot",
				Password: "iotpass",
			}).
			ApplyURI("mongodb://mongo:27017"),
	)
	log.Println("Connection to database: ")
	log.Println("Error: ", err)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("iot").Collection("testcollection")
}

func NewQueries(c *mongo.Collection) *queries.Queries {
	return queries.New(c)
}
