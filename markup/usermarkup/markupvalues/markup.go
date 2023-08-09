package markupvalues

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Default() api.InlineKeyboardMarkup {
	return api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(SmtTellingMarkupKey, SmtTellingMarkupValue),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(SmtInterestingMarkupKey, SmtInterestMarkupValue),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(SmtSleepingMarkupKey, SmtSleepingMarkupValue),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(StoriesMarkupKey, StoriesMarkupValue),
		),
	)
}
