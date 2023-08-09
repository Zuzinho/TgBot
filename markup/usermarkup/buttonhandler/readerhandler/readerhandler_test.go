package buttonhandler_test

import (
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtinteresting"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smtsleeping"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/smttelling"
	"ZuzinhoBot/markup/usermarkup/buttonhandler/story"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"testing"
)

func TestHandleReader(t *testing.T) {
	msgSlice, err := smttelling.Handle(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Smt telling")
	checkMsg(msgSlice)

	msgSlice, err = smtinteresting.Handle(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Smt interesting")
	checkMsg(msgSlice)

	msgSlice, err = smtsleeping.Handle(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Smt sleeping")
	checkMsg(msgSlice)

	msgSlice, err = story.Handle(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Story")
	checkMsg(msgSlice)
}

func checkMsg(msgSlice []*api.MessageConfig) {
	for _, msg := range msgSlice {
		log.Println(msg.Text)
	}
}
