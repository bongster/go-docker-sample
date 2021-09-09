package main

import (
	"droneia-go/src/api/db"
	"droneia-go/src/api/route"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    "database/sql"
	_ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
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

func runMigrate() {
    db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))
    if err != nil {
        panic(err)
    }
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        panic(err)
    }
    m, err := migrate.NewWithDatabaseInstance(
        "file:///app/src/db/migrations",
        "postgres", driver)
    if err != nil {
        panic(err)
    }

    if err := m.Up(); err != nil {
        m.Down()
    }
}

func main() {
    //runMigrate()
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
