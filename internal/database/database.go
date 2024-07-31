package database

import (
	"context"
	"os"
)

var (
	Client = database{}
)

type database struct {
	Client *mongo.Client
}

const (
	incidents    = "incidents"
	databaseName = "starship"
)

func Init() {
	var (
		opts *options.ClientOptions
	)
	if os.Getenv("GIN_MODE") == "" {
		opts = options.Client().ApplyUri("")
	}
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {

	}
	Client.Client = client
}
