package queries

import "go.mongodb.org/mongo-driver/mongo"

type Queries struct {
	collection *mongo.Collection
}

func New(c *mongo.Collection) *Queries {
	return &Queries{
		collection: c,
	}
}
