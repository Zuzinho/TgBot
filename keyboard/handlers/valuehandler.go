package handlers

import (
	"ZuzinhoBot/database/othertype/statistictype"
	"ZuzinhoBot/database/selector"
	"ZuzinhoBot/generator/fact"
	"ZuzinhoBot/keyboard/values"
	"ZuzinhoBot/statistic"
	"strings"
)

func UserHandler(value values.KeyboardValue) (*strings.Reader, error) {
	switch value {
	case values.FactsUserValue:
		return fact.Fact()
	case values.PhrasesUserValue, values.BoringStoriesUserValue, values.FunnyStoriesUserValue:
		return selector.SelectStory(statistictype.DataType(value))
	default:
		return nil, NewNoHandlerError(value)
	}
}

func AdminHandler(value values.KeyboardValue) (*strings.Reader, error) {
	switch value {
	case values.StatisticAdminValue:
		return statistic.Statistic(false)
	case values.StatisticSinceAdminValue:
		return statistic.Statistic(true)
	case values.StatisticPhraseAdminValue:
		return statistic.StatisticByDataType(statistictype.DataType(values.PhrasesUserValue))
	case values.StatisticFactAdminValue:
		return statistic.StatisticByDataType(statistictype.DataType(values.FactsUserValue))
	case values.StatisticBoringAdminValue:
		return statistic.StatisticByDataType(statistictype.DataType(values.BoringStoriesUserValue))
	case values.StatisticFunnyAdminValue:
		return statistic.StatisticByDataType(statistictype.DataType(values.FunnyStoriesUserValue))
	default:
		return nil, NewNoHandlerError(value)
	}
}
