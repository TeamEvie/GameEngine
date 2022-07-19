package eviecoin

import (
	"eviecoin/database"
	"eviecoin/discord"
	"github.com/disgoorg/disgo/bot"
)

type Client struct {
	Database database.Database
	Discord  bot.Client
}

var client *Client

func GetClient() *Client {
	if client != nil {
		return client
	}

	databaseInstance := database.NewDatabase()

	return &Client{
		Database: databaseInstance,
		Discord:  discord.NewClient(databaseInstance),
	}
}
