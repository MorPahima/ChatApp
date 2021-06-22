package cmd

import (
	"ChatApp/internal/data"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Module struct {
	Config Config
	Server *Server
	mongoRepo data.MongoRepository
}

func NewModule(config Config) *Module {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	mongoRepo := data.NewMongoRepository(client, context.Background())
	srv := NewServer(config, mongoRepo)

	return &Module{
		Config:    config,
		Server:    srv,
		mongoRepo: mongoRepo,
	}
}

