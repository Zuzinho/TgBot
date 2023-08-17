package handlers

import (
	"ZuzinhoBot/botapi/types"
	"ZuzinhoBot/keyboard/values"
)

func HandleUserKeyboard(value values.KeyboardValue, chatId int64) (types.Messages, error) {
	rd, err := UserHandler(value)
	if err != nil {
		return nil, err
	}

	return HandleReader(rd, chatId)
}
