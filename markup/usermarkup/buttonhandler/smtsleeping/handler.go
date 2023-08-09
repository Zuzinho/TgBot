package smtsleeping

import (
	"ZuzinhoBot/generator/sleepingstory"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/readerhandler"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler struct {
}

func (*Handler) Handle(chatId int64) ([]*api.MessageConfig, error) {
	rd, err := sleepingstory.Story()
	if err != nil {
		return nil, err
	}

	return readerhandler.HandleReader(rd, chatId)
}
