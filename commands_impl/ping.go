package commands_impl

import (
	"eviecoin/commands"
	"eviecoin/database"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var PingCommand = commands.Command{
	Name: "goping",
	CommandHandlers: map[string]commands.Handler{
		"": handlePing,
	},
}

func handlePing(_ database.Database, e *events.ApplicationCommandInteractionCreate) error {
	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetContent("pong!").
		Build(),
	)
}
