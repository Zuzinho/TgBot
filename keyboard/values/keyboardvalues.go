package keyboardvalues

type KeyboardKey string

type KeyboardValue int

const (
	PhrasesKey       KeyboardKey = "Напиши что-нибудь"
	FactsKey         KeyboardKey = "Расскажи прикольный факт"
	BoringStoriesKey KeyboardKey = "Расскажи что-нибудь заумное"
	FunnyStoriesKey  KeyboardKey = "Сборник 'Моя девушка тупень'"
)

const (
	PhrasesValue KeyboardValue = iota + 1
	FactsValue
	BoringStoriesValue
	FunnyStoriesValue
)
