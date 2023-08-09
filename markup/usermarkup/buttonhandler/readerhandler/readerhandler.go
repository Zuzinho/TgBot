package buttonhandler

import (
	"bytes"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"strings"
)

func HandleReader(rd io.Reader, chatId int64) ([]*api.MessageConfig, error) {
	msgSlice := make([]*api.MessageConfig, 0)

	buf, err := io.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	runes := bytes.Runes(buf)

	msgText := strings.Builder{}
	for _, v := range runes {
		if v == '\n' {
			newMsg := api.NewMessage(chatId, msgText.String())
			msgSlice = append(msgSlice, &newMsg)
			msgText = strings.Builder{}
			continue
		}
		msgText.WriteRune(v)
	}
	if msgText.Len() != 0 {
		newMsg := api.NewMessage(chatId, msgText.String())
		msgSlice = append(msgSlice, &newMsg)
	}

	return msgSlice, nil
}
