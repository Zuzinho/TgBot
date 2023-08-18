package values

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UserKeyboard() api.InlineKeyboardMarkup {
	return api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(PhrasesKey), string(PhrasesValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FactsKey), string(FactsValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(BoringStoriesKey), string(BoringStoriesValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FunnyStoriesKey), string(FunnyStoriesValue)),
		),
	)
}
