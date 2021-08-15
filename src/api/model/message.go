package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Chat    primitive.ObjectID `json:"chat" bson:"chat,omitempty" validate:"required"`
	Content string             `json:"content" bson:"content" validate:"required"`
}
