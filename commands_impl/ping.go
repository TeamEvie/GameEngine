package commands_impl

import (
	"eviecoin/commands"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var PingCommand = commands.Command{
	Name: "goping",
	CommandHandlers: map[string]commands.Handler{
		"": handlePing,
	},
}

func handlePing(e *events.ApplicationCommandInteractionCreate) error {
	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetContent("pong!").
		Build(),
	)
}
