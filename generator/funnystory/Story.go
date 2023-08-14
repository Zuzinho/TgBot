package funnystory

import (
	"ZuzinhoBot/database/funnystorydb"
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	story, err := funnystorydb.Story()
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(story)

	return rd, nil
}
