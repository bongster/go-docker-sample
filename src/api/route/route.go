package route

import (
	"droneia-go/src/api/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func Init(e *echo.Echo, db *mongo.Client) *echo.Echo {
	h := &handler.Handler{
		DB: db,
	}
	h.InitService()
	// Default
	e.GET("/", Index)
	e.POST("/Login", h.Login)
	e.POST("/Upload", h.UploadFile)

	e.GET("/chats", h.GetChats)
	e.POST("/chats", h.CreateChat)
	e.GET("/chats/:id/messages", h.GetChatMessages)
	e.POST("/chats/:id/messages", h.CreateChatMessages)
	e.PUT("/chats/:id", h.UpdateChat)
	e.DELETE("/chats/:id", h.DeleteChat)
	// Chatting Router

	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &handler.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/chats", h.GetChats)
	r.POST("/chats", h.CreateChat)
	r.PUT("/chats/:id", h.UpdateChat)
	r.DELETE("/chats/:id", h.DeleteChat)
	return e
}
