package commands_impl

import (
	"eviecoin/commands"
	"eviecoin/database"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"strconv"
)

var BalanceCommand = commands.Command{
	Name: "coin",
	CommandHandlers: map[string]commands.Handler{
		"balance": handleBalance,
	},
}

func handleBalance(db database.Database, e *events.ApplicationCommandInteractionCreate) error {
	user, err := db.Users.GetUser(strconv.FormatUint(uint64(e.User().ID), 10))

	if err != nil {
		return err
	}

	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetContent("You have "+strconv.Itoa(int(user.Balance))+" coins.").
		Build(),
	)
}
