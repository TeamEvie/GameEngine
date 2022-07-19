package database

import (
	"context"
	"eviecoin/environment"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v9"
	"os"
)

type Database struct {
	Users users
}

func NewDatabase() Database {
	color.Blue("[database] Getting a redis eviecoin ready...")
	getClient()

	return Database{
		Users: users{
			GetUser: getUser,
		},
	}
}

var client *redis.Client
var ctx = context.Background()

func getClient() *redis.Client {
	if client != nil {
		return client
	}

	color.Blue("[database] Creating a new redis instance...")

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       environment.GetIntWithDefault("REDIS_DB", 0),
	})

	return client
}
