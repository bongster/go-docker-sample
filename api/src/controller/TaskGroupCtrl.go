package controller

import (
	"droneia-go-api/src/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTaskGroups(c echo.Context) error {
	db, err := NewDB("postgres://droneina:droneina@192.168.1.9:15432/droneina")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	sqlStatement := `SELECT id, name, path, status from "TaskGroups" order by id`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := model.TaskGroups{}
	for rows.Next() {
		taskGroup := model.TaskGroup{}
		err2 := rows.Scan(&taskGroup.Id, &taskGroup.Name, &taskGroup.Path, &taskGroup.Status)
		if err2 != nil {
			return err2
		}
		result.TaskGroups = append(result.TaskGroups, taskGroup)
	}
	return c.JSON(http.StatusOK, result)
}
