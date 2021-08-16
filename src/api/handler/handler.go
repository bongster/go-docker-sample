package handler

import (
	"go.mongodb.org/mongo-driver/mongo"

	service "droneia-go/src/api/service"
)

type Handler struct {
	DB          *mongo.Client
	ChatService service.ChatService
}

func (h *Handler) InitService() {
	h.ChatService = service.ChatService{
		DB: h.DB,
	}
}
