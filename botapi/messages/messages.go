package messages

import (
	"ZuzinhoBot/botapi/types"
	"ZuzinhoBot/keyboard/values"
	"ZuzinhoBot/users"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func FirstMessage(role users.Role, chatId int64) (types.Message, error) {
	switch role {
	case users.Admin:
		msg := api.NewMessage(chatId, "Здраствуйте, хозяин!\nЧего желаете?")
		msg.ReplyMarkup = values.AdminKeyboard()

		return &msg, nil
	case users.User:
		msg := api.NewMessage(chatId, "Привет!!\nЧего надобно от меня?")
		msg.ReplyMarkup = values.UserKeyboard()

		return &msg, nil
	case users.UnwantedUser:
		msg := api.NewMessage(chatId, "Ты кто?\nЯ тебя не знаю\nНичего тебе делать не буду")

		return &msg, nil
	default:
		return nil, users.NewUnknownRoleError(role)
	}
}

func HelpMessage(role users.Role, chatId int64) (types.Message, error) {
	switch role {
	case users.User, users.Admin:
		msg := api.NewMessage(chatId, "Нажми на любую кнопку")

		return &msg, nil
	case users.UnwantedUser:
		msg := api.NewMessage(chatId, "Я тебе сказал, я ничего делать не буду\nХоть досметри пиши мне")

		return &msg, nil
	default:
		return nil, users.NewUnknownRoleError(role)
	}

}

func LastMessage(role users.Role, chatId int64) (types.Message, error) {
	switch role {
	case users.Admin:
		msg := api.NewMessage(chatId, "Хотите еще, господин?")
		msg.ReplyMarkup = values.AdminKeyboard()

		return &msg, nil
	case users.User:
		msg := api.NewMessage(chatId, "Еще?")
		msg.ReplyMarkup = values.UserKeyboard()

		return &msg, nil
	default:
		return nil, users.NewUnknownRoleError(role)
	}
}
