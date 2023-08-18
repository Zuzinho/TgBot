package values

type KeyboardKey string

type KeyboardValue string

const (
	PhrasesUserKey       KeyboardKey = "Напиши что-нибудь"
	FactsUserKey         KeyboardKey = "Расскажи прикольный факт"
	BoringStoriesUserKey KeyboardKey = "Расскажи что-нибудь заумное"
	FunnyStoriesUserKey  KeyboardKey = "Сборник 'Моя девушка тупень'"
)

const (
	PhrasesUserValue       KeyboardValue = "Фраза"
	FactsUserValue         KeyboardValue = "Факт"
	BoringStoriesUserValue KeyboardValue = "Заумное"
	FunnyStoriesUserValue  KeyboardValue = "История"
)
