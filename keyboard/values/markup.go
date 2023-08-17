package keyboardvalues

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func Default() api.InlineKeyboardMarkup {
	return api.NewInlineKeyboardMarkup(
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(PhrasesKey), strconv.Itoa(int(PhrasesValue))),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FactsKey), strconv.Itoa(int(FactsValue))),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(BoringStoriesKey), strconv.Itoa(int(BoringStoriesValue))),
		),
		api.NewInlineKeyboardRow(
			api.NewInlineKeyboardButtonData(string(FunnyStoriesKey), strconv.Itoa(int(FunnyStoriesValue))),
		),
	)
}
