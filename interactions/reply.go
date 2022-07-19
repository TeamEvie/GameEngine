package interactions

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func Defer(e *events.ApplicationCommandInteractionCreate, ephemeral bool) {
	err := e.DeferCreateMessage(ephemeral)
	if err != nil {
		return
	}
}

func EditReply(e *events.ApplicationCommandInteractionCreate, update discord.MessageUpdate) (*discord.Message, error) {
	response, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), update)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func VoidEditReply(e *events.ApplicationCommandInteractionCreate, update discord.MessageUpdate) error {
	_, err := EditReply(e, update)
	return err
}
