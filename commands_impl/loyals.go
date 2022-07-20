package commands_impl

import (
	"eviecoin/commands"
	"eviecoin/database"
	"eviecoin/interactions"
	"eviecoin/responses"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"strconv"
	"time"
)

var LoyalCommands = commands.Command{
	Name: "loyals",
	CommandHandlers: map[string]commands.Handler{
		"daily": handleDaily,
	},
}

func handleDaily(db database.Database, e *events.ApplicationCommandInteractionCreate) error {
	interactions.Defer(e, false)

	user, err := db.Users.GetUser(strconv.FormatUint(uint64(e.User().ID), 10))

	if err != nil {
		return err
	}

	// make sure the user hasn't redeemed more than once every 24 hours
	if time.Since(user.LastRedeemedDaily) < time.Hour*24 {
		return interactions.VoidEditReply(e, discord.NewMessageUpdateBuilder().SetEmbeds(responses.NewErrorResponse("You have already redeemed your daily coins!").Build()).Build())
	}

	// add the daily coins to the user
	user, err = db.Users.AddBalance(strconv.FormatUint(uint64(e.User().ID), 10), 100, database.NewSummonFrom("daily"))
	if err != nil {
		return err
	}

	// update the user's last redeemed daily
	user.LastRedeemedDaily = time.Now()

	user, err = db.Users.SetUser(strconv.FormatUint(uint64(e.User().ID), 10), user)
	if err != nil {
		return err
	}

	return interactions.VoidEditReply(e, discord.NewMessageUpdateBuilder().
		SetEmbeds(responses.NewBankResponse(e.User().Username, user, "You have redeemed your daily coins!")).
		Build(),
	)

}
