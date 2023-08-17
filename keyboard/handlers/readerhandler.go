package handlers

import (
	"ZuzinhoBot/botapi/types"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"strings"
)

func HandleReader(reader *strings.Reader, chatId int64) (types.Messages, error) {
	msgSlice := make(types.Messages, 0)

	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	split := strings.Split(string(buf), "\n")
	for _, text := range split {
		if text == "" {
			continue
		}
		msg := api.NewMessage(chatId, text)
		msgSlice = append(msgSlice, &msg)
	}

	return msgSlice, nil
}
