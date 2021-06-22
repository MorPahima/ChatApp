package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongoRepository struct {
	client *mongo.Client
	ctx    context.Context
}

func NewMongoRepository(client *mongo.Client, ctx context.Context) *mongoRepository {
	return &mongoRepository{client: client, ctx: ctx}
}

func (this *mongoRepository) Set(s string) error {

	user := User{Name: s}

	collection := this.client.Database("ChatApplication").Collection("users")
	result, err := collection.InsertOne(this.ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return nil
}

type User struct {
	Name string `bson:"name"`
}