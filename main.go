package main

import (
	"ZuzinhoBot/botapi/messages"
	"ZuzinhoBot/database/insertor"
	"ZuzinhoBot/database/othertype/statistictype"
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

	isStarted := false

	errlog.LogOnErr(err)

	for update := range updates {
		if msg := update.Message; msg != nil {
			var role users.Role
			userName := msg.Chat.UserName
			log.Printf("User '%s' sent message '%s'", userName, msg.Text)
			if chatId := msg.Chat.ID; msg.Text == "/start" {
				err = insertor.InsertUnwantedUser(chatId, userName)
				errlog.LogOnErr(err)
				role = users.UnwantedUser
			} else {
				role, err = users.NewRole(chatId, userName)
				if err != nil {
					log.Println(err)
					continue
				}
			}

			log.Printf("User '%s' has role '%s'", userName, role)

			if chatId := msg.Chat.ID; !isStarted {
				firstMsg, err := messages.FirstMessage(role, chatId)
				if err != nil {
					log.Println(err)
					continue
				}

				_, err = bot.Send(*firstMsg)
				errlog.LogOnErr(err)
				isStarted = true
			} else {
				helpMsg, err := messages.HelpMessage(role, chatId)
				if err != nil {
					log.Println(err)
					continue
				}

				_, err = bot.Send(*helpMsg)
				errlog.LogOnErr(err)
			}

			continue
		}

		if query := update.CallbackQuery; query != nil {
			userName := query.Message.Chat.UserName
			chatId := query.Message.Chat.ID
			value := values.KeyboardValue(query.Data)

			callback := api.NewCallback(query.ID, query.Data)
			_, err = bot.AnswerCallbackQuery(callback)
			errlog.LogOnErr(err)

			log.Printf("User '%s' sent query with data '%s'", userName, value)

			role, err := users.NewRole(chatId, userName)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("User '%s' has role '%s'", userName, role)

			msgSlice, err := handlers.HandleKeyboard(value, chatId, role)
			errlog.LogOnErr(err)

			for _, msg := range msgSlice {
				log.Printf("Sent message '%s' to user '%s'", msg.Text, userName)
				_, err = bot.Send(*msg)
				errlog.LogOnErr(err)
				time.Sleep(1500 * time.Millisecond)
			}

			lastMsg, err := messages.LastMessage(role, chatId)
			if err != nil {
				log.Println(err)
				continue
			}

			_, err = bot.Send(*lastMsg)
			errlog.LogOnErr(err)

			err = insertor.InsertUserMessage(chatId, statistictype.DataType(value), time.Now())
			errlog.LogOnErr(err)
			continue
		}
	}
}
