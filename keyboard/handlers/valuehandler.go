package handlers

import (
	"ZuzinhoBot/generator/boringstory"
	"ZuzinhoBot/generator/fact"
	"ZuzinhoBot/generator/funnystory"
	"ZuzinhoBot/generator/phrase"
	"ZuzinhoBot/keyboard/values"
	"strings"
)

func UserHandler(value values.KeyboardValue) (*strings.Reader, error) {
	switch value {
	case values.PhrasesUserValue:
		return phrase.Phrase()
	case values.FactsUserValue:
		return fact.Fact()
	case values.BoringStoriesUserValue:
		return boringstory.Story()
	case values.FunnyStoriesUserValue:
		return funnystory.Story()
	default:
		return nil, NewNoHandlerError(value)
	}
}
