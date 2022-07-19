package discord

import (
	"context"
	"eviecoin/commands"
	"eviecoin/database"
	"eviecoin/environment"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/fatih/color"
)

func NewClient(db database.Database) bot.Client {
	commandManager := commands.NewManager(db)

	LoadCommands(commandManager)

	client, err := disgo.New(environment.GetStringNullPanic("DISCORD_TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNone,
			),
		),
		bot.WithEventListenerFunc(commandManager.OnApplicationCommandInteraction),
	)

	if err != nil {
		panic(err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	color.Green("[discord] Connected to the gateway!")

	return client
}
