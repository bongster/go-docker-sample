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
	e.GET("/TaskGroups", controller.GetTaskGroups)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
