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

const (
	StatisticAdminKey       KeyboardKey = "Получить общую статистику"
	StatisticSinceAdminKey  KeyboardKey = "Получить статистику с последнего раза"
	StatisticPhraseAdminKey KeyboardKey = "Статистика 'Фраза'"
	StatisticFactAdminKey   KeyboardKey = "Статистика 'Факт'"
	StatisticBoringAdminKey KeyboardKey = "Статистика 'Заумное'"
	StatisticFunnyAdminKey  KeyboardKey = "Статистика 'История'"
)

const (
	StatisticAdminValue       KeyboardValue = "Статистика"
	StatisticSinceAdminValue  KeyboardValue = "Последняя статистика"
	StatisticPhraseAdminValue KeyboardValue = "Статистика 'Фраза'"
	StatisticFactAdminValue   KeyboardValue = "Статистика 'Факт'"
	StatisticBoringAdminValue KeyboardValue = "Статистика 'Заумное'"
	StatisticFunnyAdminValue  KeyboardValue = "Статистика 'История'"
)
