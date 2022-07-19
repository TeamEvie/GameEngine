package responses

import (
	"github.com/disgoorg/disgo/discord"
	"time"
)

func NewInfoResponse() *discord.EmbedBuilder {
	return discord.NewEmbedBuilder().
		SetColor(0x2f3136).
		SetTimestamp(time.Now()).
		SetFooterText("$EVIE")
}
