package database

import (
	"eviecoin/environment"
	"github.com/go-redis/redis/v9"
	"os"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       environment.GetIntWithDefault("REDIS_DB", 0),
	})
}
