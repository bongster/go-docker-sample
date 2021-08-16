package service

import (
	"context"
	"droneia-go/src/api/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IChatService interface {
	FindAll(options *options.FindOptions) ([]*model.Chat, error)
	InsertOne(data *model.Chat) (*model.Chat, error)
}

type ChatService struct {
	DB *mongo.Client
}

func (c *ChatService) FindAll(options *options.FindOptions) ([]*model.Chat, error) {
	collection := c.DB.Database("app").Collection("chats")
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

func (c *ChatService) InsertOne(data *model.Chat) (*model.Chat, error) {
	collection := c.DB.Database("app").Collection("chats")
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var chat *model.Chat
	err = collection.FindOne(context.TODO(), bson.M{"_id": insertResult.InsertedID}).Decode(&chat)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return chat, nil
}
