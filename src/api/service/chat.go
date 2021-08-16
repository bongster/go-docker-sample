package service

import (
	"context"
	"droneia-go/src/api/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatService struct {
	DB *mongo.Client
}

func (c *ChatService) getCollection() *mongo.Collection {
	collection := c.DB.Database("app").Collection("chats")
	return collection
}

func (c *ChatService) FindAll(options *options.FindOptions) ([]*model.Chat, error) {
	collection := c.getCollection()
	var results []*model.Chat
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem model.Chat
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}
	defer cur.Close(context.TODO())
	return results, nil
}
