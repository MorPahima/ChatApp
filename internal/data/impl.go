package data

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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


func (this *mongoRepository) GetAll() ([]interface{}, error){
	var data []interface{}
	collection := this.client.Database("ChatApplication").Collection("users")

	cur, err := collection.Find(this.ctx, bson.D{{}})
	if err != nil {
		logrus.Infof("could not pull all the data in DB")
		return nil, err
	}

	for cur.Next(this.ctx) {
		var t interface{}
		err = cur.Decode(&t)
		if err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	_ = cur.Close(this.ctx)

	if len(data) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return data, nil
}