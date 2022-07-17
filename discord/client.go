package discord

import (
	"context"
	"eviecoin/commands"
	"eviecoin/environment"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
)

func NewClient() bot.Client {
	commandManager := commands.NewManager()

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

	return client
}
