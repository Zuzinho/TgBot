package phrase

import (
	"ZuzinhoBot/database/phrasedb"
	"io"
	"strings"
)

func Phrase() (io.Reader, error) {
	phrase, err := phrasedb.Phrase()
	if err != nil {
		return nil, err
	}
	
	rd := strings.NewReader(phrase)

	return rd, nil
}
