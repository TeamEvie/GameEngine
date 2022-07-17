package client

import (
	"eviecoin/database"
	"eviecoin/discord"
	"github.com/disgoorg/disgo/bot"
	"github.com/go-redis/redis/v9"
)

type Client struct {
	redis   *redis.Client
	discord bot.Client
}

var client *Client

func GetClient() *Client {
	if client != nil {
		return client
	}

	return &Client{
		redis:   database.NewClient(),
		discord: discord.NewClient(),
	}
}
