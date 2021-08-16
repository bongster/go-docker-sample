package main

import (
	"droneia-go/src/api/db"
	"droneia-go/src/api/route"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	_, err := db.NewDB(os.Getenv(("DB_URL")))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("DB Connected")

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db, _ := db.NewMongoDB(os.Getenv("MONGO_DB_URL"))

	e = route.Init(e, db)

	if value, ok := os.LookupEnv("PORT"); ok {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", value)))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
