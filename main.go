package main

import (
	"ZuzinhoBot/env"
	"ZuzinhoBot/errlog"
	"ZuzinhoBot/keyboard/handlers"
	"ZuzinhoBot/keyboard/values"
	"ZuzinhoBot/users"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func main() {
	bot, err := api.NewBotAPI(env.Token)
	errlog.LogOnErr(err)

	log.Printf("Authorized bot %s\n", bot.Self.UserName)

	u := api.NewUpdate(0)

	updates, err := bot.GetUpdatesChan(u)

	isStart := false

	errlog.LogOnErr(err)

	for update := range updates {
		if update.Message != nil {
			sentMsg := update.Message
			firstName := sentMsg.Chat.FirstName
			log.Printf("User %s sent message with text '%s'", firstName, sentMsg.Text)
			switch users.GetRole(firstName) {
			case users.User:
				log.Printf("User %s has role 'User'", firstName)
				if !isStart {
					msg := api.NewMessage(sentMsg.Chat.ID, "Привет")
					msg.ReplyMarkup = values.UserKeyboard()

					_, err = bot.Send(msg)
					errlog.LogOnErr(err)
					isStart = true
				} else {
					msg := api.NewMessage(sentMsg.Chat.ID, "Нажми на какую-нибудь кнопку")

					_, err = bot.Send(msg)
					errlog.LogOnErr(err)
				}
			case users.Admin:
				log.Printf("User %s has role 'Admin'", firstName)
				msg := api.NewMessage(sentMsg.Chat.ID, "Здраствуйте, хозяин")
				_, err = bot.Send(msg)
				errlog.LogOnErr(err)
			}
		} else if update.CallbackQuery != nil {
			sentQuery := update.CallbackQuery
			firstName := sentQuery.Message.Chat.FirstName
			log.Printf("User %s sent query with data '%s'", firstName, sentQuery.Data)
			switch users.GetRole(sentQuery.Message.Chat.FirstName) {
			case users.User:
				log.Printf("User %s has role 'User'", firstName)

				callback := api.NewCallback(sentQuery.ID, sentQuery.Data)
				_, err := bot.AnswerCallbackQuery(callback)
				errlog.LogOnErr(err)
				msgSlice, err := handlers.HandleUserKeyboard(values.KeyboardValue(sentQuery.Data), sentQuery.Message.Chat.ID)
				errlog.LogOnErr(err)

				for _, msg := range msgSlice {
					_, err = bot.Send(*msg)
					errlog.LogOnErr(err)
					log.Printf("Message '%s' was sent to user %s", msg.Text, firstName)
					time.Sleep(time.Second)
				}
				lastMsg := api.NewMessage(sentQuery.Message.Chat.ID, "Еще?")
				lastMsg.ReplyMarkup = values.UserKeyboard()
				_, err = bot.Send(lastMsg)
				errlog.LogOnErr(err)
			}
		}
	}
}
