package usermarkup

import (
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtinteresting"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtsleeping"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smttelling"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/story"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDefault(query *api.CallbackQuery, bot *api.BotAPI) ([]*api.MessageConfig, error) {
	callback := api.NewCallback(query.ID, query.Data)
	_, err := bot.AnswerCallbackQuery(callback)
	if err != nil {
		return nil, err
	}

	var msgSlice []*api.MessageConfig
	chatId := query.Message.Chat.ID

	switch query.Data {
	case SmtTellingMarkupValue:
		msgSlice, err = smttelling.Handle(chatId)
	case SmtInterestMarkupValue:
		msgSlice, err = smtinteresting.Handle(chatId)
	case SmtSleepingMarkupValue:
		msgSlice, err = smtsleeping.Handle(chatId)
	case StoriesMarkupValue:
		msgSlice, err = story.Handle(chatId)
	default:
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return msgSlice, nil
}
