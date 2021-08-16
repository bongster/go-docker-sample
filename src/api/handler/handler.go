package handler

import (
	"go.mongodb.org/mongo-driver/mongo"

	service "droneia-go/src/api/service"
)

type Handler struct {
	DB          *mongo.Client
	ChatService service.IChatService
}

func (h *Handler) InitService() {
	if h.ChatService == nil {
		h.ChatService = &service.ChatService{
			DB: h.DB,
		}
	}

}
