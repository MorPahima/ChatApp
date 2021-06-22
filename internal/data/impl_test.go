package data

import (
	"ChatApp/cmd"
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongoRepository_Set(t *testing.T) {
	cfg, err := cmd.ParseConfig()
	assert.Nilf(t, err, "err %v", err)

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	assert.Nilf(t, err, "err %v", err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	assert.Nilf(t, err, "err %v", err)

	defer client.Disconnect(ctx)

	mongoRepo := NewMongoRepository(client, ctx)
	err = mongoRepo.Set("Yay it's working :D")
	assert.Nilf(t, err, "err %v", err)
}


