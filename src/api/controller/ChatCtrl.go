package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetChats(c echo.Context) error {
	db, err := NewDB("postgres://droneina:droneina@192.168.1.9:15432/droneina")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()
	dummies := make([]interface{}, 0)
	return c.JSON(http.StatusOK, dummies)
}

func GetChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusOK, dummy)
}

func CreateChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func UpdateChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}

func DeleteChat(c echo.Context) error {
	dummy := new(interface{})
	return c.JSON(http.StatusNotImplemented, dummy)
}
