package usermarkup

import (
	"ZuzinhoBot/markup/usermarkup/buttonhandler"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDefault(query *api.CallbackQuery, bot *api.BotAPI) ([]*api.MessageConfig, error) {
	callback := api.NewCallback(query.ID, query.Data)
	_, err := bot.AnswerCallbackQuery(callback)
	if err != nil {
		return nil, err
	}

	chatId := query.Message.Chat.ID

	handler, err := buttonhandler.GetHandler(query.Data)
	if err != nil {
		return nil, err
	}

	msgSlice, err := handler.Handle(chatId)
	if err != nil {
		return nil, err
	}

	return msgSlice, nil
}
