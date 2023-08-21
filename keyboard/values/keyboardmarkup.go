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

func AdminKeyboard() api.InlineKeyboardMarkup {
	return api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(StatisticAdminKey), string(StatisticAdminValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(StatisticSinceAdminKey), string(StatisticSinceAdminValue)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(StatisticPhraseAdminKey), string(StatisticPhraseAdminValue)),
			api.NewInlineKeyboardButtonData(string(StatisticFactAdminKey), string(StatisticFactAdminKey)),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(StatisticBoringAdminKey), string(StatisticBoringAdminValue)),
			api.NewInlineKeyboardButtonData(string(StatisticFunnyAdminKey), string(StatisticFunnyAdminValue)),
		),
	)
}
