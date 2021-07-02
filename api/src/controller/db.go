package controller

import (
	"database/sql"
	"fmt"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("%s?sslmode=disable", dataSourceName))
	if err != nil {
		return nil, err
	}
	if err1 := db.Ping(); err1 != nil {
		return nil, err1
	}
	return db, nil
}
