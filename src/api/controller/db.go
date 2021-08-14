package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func NewMongoDB(dataSourceName string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(dataSourceName)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
