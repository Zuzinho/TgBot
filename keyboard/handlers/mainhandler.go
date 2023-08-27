package handlers

import (
	"ZuzinhoBot/botapi/types"
	"ZuzinhoBot/keyboard/values"
	"ZuzinhoBot/users"
)

func HandleUserKeyboard(value values.KeyboardValue, chatId int64) (types.Messages, error) {
	rd, err := UserHandler(value)
	if err != nil {
		return nil, err
	}

	return HandleReader(rd, chatId)
}

func HandleAdminKeyboard(value values.KeyboardValue, chatId int64) (types.Messages, error) {
	rd, err := AdminHandler(value)
	if err != nil {
		return nil, err
	}

	return HandleReader(rd, chatId)
}

func HandleKeyboard(value values.KeyboardValue, chatId int64, role users.Role) (types.Messages, error) {
	switch role {
	case users.Admin:
		return HandleAdminKeyboard(value, chatId)
	case users.User:
		return HandleUserKeyboard(value, chatId)
	default:
		return nil, NewInAccessedHandlerError(value)
	}
}
