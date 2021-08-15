package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "droneia-go/src/api/model"
)

func GetChats(c echo.Context) error {
	client, err := NewMongoDB("mongodb://admin:admin@mongo:27017")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	collection := client.Database("app").Collection("chats")
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []*model.Chat
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for cur.Next(context.TODO()) {
		var elem model.Chat
		err := cur.Decode(&elem)
		if err != nil {
			return err
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return err
	}
	defer cur.Close(context.TODO())

	return c.JSON(http.StatusOK, results)
}

func GetChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusOK, dummy)
}

func CreateChat(c echo.Context) error {
	client, err := NewMongoDB("mongodb://admin:admin@mongo:27017")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	model := new(model.Chat)
	if err = c.Bind(model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(model); err != nil {
		return err
	}
	fmt.Printf("%v", model)
	collection := client.Database("app").Collection("chats")
	insertResult, err := collection.InsertOne(context.TODO(), model)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusCreated, insertResult)
}

func UpdateChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func DeleteChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func GetChatMessages(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	client, err := NewMongoDB("mongodb://admin:admin@mongo:27017")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	collection := client.Database("app").Collection("messages")
	var results []*model.Message
	filterStage := bson.D{primitive.E{Key: "$limit", Value: 10}}
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "chat", Value: id}}}}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, filterStage})
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return err
	}

	defer cur.Close(context.TODO())

	return c.JSON(http.StatusOK, results)
}

func CreateChatMessages(c echo.Context) error {
	id := c.Param("id")
	client, err := NewMongoDB("mongodb://admin:admin@mongo:27017")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	m := new(model.Message)
	m.Chat, _ = primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(m); err != nil {
		return err
	}
	collection := client.Database("app").Collection("messages")
	insertResult, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusCreated, insertResult)
}
