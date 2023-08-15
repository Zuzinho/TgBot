package funnystory

import (
	"ZuzinhoBot/database"
	"ZuzinhoBot/database/selector"
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	story, err := selector.Select(database.FunnyStoriesTableName)
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(story)

	return rd, nil
}
