package buttonhandler

import (
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtinteresting"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtsleeping"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smttelling"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/story"
	"ZuzinhoBot/markup/usermarkup/markupvalues"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler interface {
	Handle(chatId int64) ([]*api.MessageConfig, error)
}

func GetHandler(btnCnt string) (Handler, error) {
	switch btnCnt {
	case markupvalues.SmtTellingMarkupValue:
		hnd := smttelling.Handler{}
		return &hnd, nil
	case markupvalues.SmtInterestMarkupValue:
		hnd := smtinteresting.Handler{}
		return &hnd, nil
	case markupvalues.SmtSleepingMarkupValue:
		hnd := smtsleeping.Handler{}
		return &hnd, nil
	case markupvalues.StoriesMarkupValue:
		hnd := story.Handler{}
		return &hnd, nil
	default:
		return nil, NewNoHandlerError(btnCnt)
	}
}
