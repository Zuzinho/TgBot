package values

type KeyboardKey string

type KeyboardValue string

const (
	PhrasesKey       KeyboardKey = "Напиши что-нибудь"
	FactsKey         KeyboardKey = "Расскажи прикольный факт"
	BoringStoriesKey KeyboardKey = "Расскажи что-нибудь заумное"
	FunnyStoriesKey  KeyboardKey = "Сборник 'Моя девушка тупень'"
)

const (
	PhrasesValue KeyboardValue = "Фраза"
	FactsValue KeyboardValue = "Факт"
	BoringStoriesValue KeyboardValue = "Заумное"
	FunnyStoriesValue KeyboardValue = "История"
)

