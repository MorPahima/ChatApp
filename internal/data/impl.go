package data

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	client *mongo.Client
	ctx    context.Context
}

func NewMongoRepository(client *mongo.Client, ctx context.Context) *mongoRepository {
	return &mongoRepository{client: client, ctx: ctx}
}

func (this *mongoRepository) Set(body interface{}) error {

	collection := this.client.Database("ChatApplication").Collection("users")
	_, err := collection.InsertOne(this.ctx, body)
	if err != nil {
		logrus.Infof("could not insert body %+v", body)
		return err
	}
	return nil
}