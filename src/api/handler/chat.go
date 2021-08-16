package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "droneia-go/src/api/model"
)

func (h *Handler) GetChats(c echo.Context) error {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	results, err := h.ChatService.FindAll(findOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, results)
}

// change use Handler instead of class method for testing
func (h *Handler) GetChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusOK, dummy)
}

func (h *Handler) CreateChat(c echo.Context) error {
	model := new(model.Chat)
	if err := c.Bind(model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(model); err != nil {
		return err
	}
	collection := h.DB.Database("app").Collection("chats")
	insertResult, err := collection.InsertOne(context.TODO(), model)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusCreated, insertResult)
}

func (h *Handler) UpdateChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func (h *Handler) DeleteChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func (h *Handler) GetChatMessages(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	collection := h.DB.Database("app").Collection("messages")
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

func (h *Handler) CreateChatMessages(c echo.Context) error {
	id := c.Param("id")

	m := new(model.Message)
	m.Chat, _ = primitive.ObjectIDFromHex(id)

	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(m); err != nil {
		return err
	}
	collection := h.DB.Database("app").Collection("messages")
	insertResult, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusCreated, insertResult)
}
