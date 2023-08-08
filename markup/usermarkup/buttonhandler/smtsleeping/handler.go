package smtsleeping

import (
	"ZuzinhoBot/generator/sleepingstory"
	"ZuzinhoBot/markup/usermarkup/buttonhandler"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Handle(chatId int64) ([]*api.MessageConfig, error) {
	rd, err := sleepingstory.Story()
	if err != nil {
		return nil, err
	}

	return buttonhandler.HandleReader(rd, chatId)
}
