package usermarkup

import (
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtinteresting"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtsleeping"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smttelling"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/story"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDefault(update api.Update, bot *api.BotAPI) ([]*api.MessageConfig, error) {
	callback := api.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	_, err := bot.AnswerCallbackQuery(callback)
	if err != nil {
		return nil, err
	}

	var msgSlice []*api.MessageConfig
	chatId := update.CallbackQuery.Message.Chat.ID

	switch update.CallbackQuery.Data {
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
