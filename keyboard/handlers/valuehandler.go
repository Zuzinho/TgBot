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
	case values.PhrasesValue:
		return phrase.Phrase()
	case values.FactsValue:
		return fact.Fact()
	case values.BoringStoriesValue:
		return boringstory.Story()
	case values.FunnyStoriesValue:
		return funnystory.Story()
	default:
		return nil, NewNoHandlerError(value)
	}
}
