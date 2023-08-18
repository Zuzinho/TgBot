package values

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UserKeyboard() api.InlineKeyboardMarkup {
	return api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(PhrasesUserKey), string(PhrasesUserValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FactsUserKey), string(FactsUserValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(BoringStoriesUserKey), string(BoringStoriesUserValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FunnyStoriesUserKey), string(FunnyStoriesUserValue)),
		),
	)
}
