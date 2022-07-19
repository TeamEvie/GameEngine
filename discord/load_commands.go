package discord

import (
	"eviecoin/commands"
	"eviecoin/commands_impl"
)

func LoadCommands(m *commands.Manager) {
	m.LoadCommands(
		commands_impl.PingCommand,
		commands_impl.BalanceCommand,
	)
}
