package sleepingstory

import (
	"ZuzinhoBot/database/sleepingstorydb"
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	story, err := sleepingstorydb.Story()
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(story)

	return rd, nil
}
