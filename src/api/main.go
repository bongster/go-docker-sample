package main

import (
	"droneia-go/src/api/controller"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_, err := controller.NewDB(os.Getenv(("DB_URL")))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("DB Connected")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.POST("/Login", controller.Login)
	e.POST("/Upload", controller.UploadFile)
	// Chatting Router

	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &controller.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/chats", controller.GetChats)
	r.POST("/chats", controller.CreateChat)
	r.PUT("/chats/:id", controller.UpdateChat)
	r.DELETE("/chats/:id", controller.DeleteChat)

	if value, ok := os.LookupEnv("PORT"); ok {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", value)))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
