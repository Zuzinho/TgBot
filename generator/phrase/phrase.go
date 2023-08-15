package phrase

import (
	"ZuzinhoBot/database"
	"ZuzinhoBot/database/selector"
	"io"
	"strings"
)

func Phrase() (io.Reader, error) {
	phrase, err := selector.Select(database.PhrasesTableName)
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(phrase)

	return rd, nil
}
