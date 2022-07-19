package commands_impl

import (
	"eviecoin/commands"
	"eviecoin/database"
	"eviecoin/interactions"
	"eviecoin/responses"
	"fmt"
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
	interactions.Defer(e, false)

	user, err := db.Users.GetUser(strconv.FormatUint(uint64(e.User().ID), 10))

	if err != nil {
		return err
	}

	return interactions.VoidEditReply(e, discord.NewMessageUpdateBuilder().
		SetEmbeds(responses.NewInfoResponse().
			SetTitle("Balance").
			SetDescription(fmt.Sprintf("%d coins", user.Balance)).
			Build()).
		Build(),
	)
}
