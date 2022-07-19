package commands

import (
	"eviecoin/database"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/fatih/color"
)

func (manager *Manager) LoadCommands(commands ...Command) {
	for _, command := range commands {
		manager.Commands[command.Name] = command
		color.Yellow("[commands] Loaded command: %s", command.Name)
	}
}

func (manager *Manager) OnApplicationCommandInteraction(e *events.ApplicationCommandInteractionCreate) {
	if command, ok := manager.Commands[e.Data.CommandName()]; ok {
		var path string

		if data, ok := e.Data.(discord.SlashCommandInteractionData); ok {
			if data.SubCommandGroupName != nil {
				path += *data.SubCommandGroupName + "/"
			}
			if data.SubCommandName != nil {
				path += *data.SubCommandName
			}
		}

		if handler, ok := command.CommandHandlers[path]; ok {
			if err := handler(manager.Database, e); err != nil {
				color.Red("Error handling command: %s", err)
			}
			return
		}
		return
	}
}

type (
	Handler func(db database.Database, e *events.ApplicationCommandInteractionCreate) error
	Command struct {
		Name            string
		CommandHandlers map[string]Handler
	}
)
