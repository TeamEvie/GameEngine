package responses

import (
	"github.com/disgoorg/disgo/discord"
)

func NewInfoResponse() *discord.EmbedBuilder {
	return discord.NewEmbedBuilder().
		SetColor(0x2f3136)
}
