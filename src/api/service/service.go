package service

import (
	"droneia-go/src/api/model"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	FindAll(options *options.FindOptions) ([]*model.Chat, error)
}
