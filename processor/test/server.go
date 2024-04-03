package test

import (
	"iotvisual/processor/internal/queries"
	"iotvisual/processor/internal/server"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewServer() *server.Server {
	s := server.New(
		server.WithLogger(),
		server.WithDatabase(),
	)
	return s
}

func NewQueries(c *mongo.Collection) *queries.Queries {
	return queries.New(c)
}
