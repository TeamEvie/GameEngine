package responses

import (
	"eviecoin/database"
	"eviecoin/responses/emojis"
	"fmt"
	"github.com/disgoorg/disgo/discord"
)

func NewBankResponse(username string, user database.User, extraInfo ...string) discord.Embed {
	extraInfoString := ""
	for _, info := range extraInfo {
		extraInfoString += info + " "
	}
	extraInfoString = "> " + extraInfoString + "\n"

	return NewInfoResponse().
		SetTitle(fmt.Sprintf("%s's balance", username)).
		SetDescription(fmt.Sprintf("%s**$EVIE**: %s%d", extraInfoString, emojis.EvieCoin, user.Balance)).
		Build()
}
