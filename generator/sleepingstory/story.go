package sleepingstory

import (
	"ZuzinhoBot/database/selector"
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	story, err := selector.Select("sleeping_stories")
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(story)

	return rd, nil
}
