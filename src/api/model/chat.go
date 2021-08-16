package model

type Chat struct {
	Id     interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string      `json:"name" validate:"required"`
	Status string      `json:"status" validate:"required"`
}
