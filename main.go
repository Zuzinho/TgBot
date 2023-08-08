package main

import (
	"ZuzinhoBot/errlog"
	"ZuzinhoBot/markup/usermarkup"
	"ZuzinhoBot/token"
	"ZuzinhoBot/users"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func main() {
	bot, err := api.NewBotAPI(token.Token())
	errlog.LogOnErr(err)

	bot.Debug = true

	log.Printf("Authorized bot %s\n", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	isStart := false

	errlog.LogOnErr(err)

	for update := range updates {
		if update.Message != nil {
			switch users.GetRole(update.Message.Chat.FirstName) {
			case users.User:
				if !isStart {
					msg := api.NewMessage(update.Message.Chat.ID, "Привет")
					msg.ReplyMarkup = usermarkup.Default()

					_, err = bot.Send(msg)
					errlog.LogOnErr(err)
					isStart = true
				} else {
					msg := api.NewMessage(update.Message.Chat.ID, "Нажми на какую-нибудь кнопку")

					_, err = bot.Send(msg)
					errlog.LogOnErr(err)
				}
			case users.Admin:
				msg := api.NewMessage(update.Message.Chat.ID, "Здраствуйте, хозяин")
				_, err = bot.Send(msg)
				errlog.LogOnErr(err)
			}
		} else if update.CallbackQuery != nil {
			switch users.GetRole(update.CallbackQuery.Message.Chat.FirstName) {
			case users.User:
				msgSlice, err := usermarkup.HandleDefault(update, bot)
				errlog.LogOnErr(err)

				for _, msg := range msgSlice {
					_, err = bot.Send(*msg)
					errlog.LogOnErr(err)
					time.Sleep(time.Second)
				}
				lastMsg := api.NewMessage(update.CallbackQuery.Message.Chat.ID, "Еще?")
				lastMsg.ReplyMarkup = usermarkup.Default()
				_, err = bot.Send(lastMsg)
				errlog.LogOnErr(err)
			}
		}
	}
}
