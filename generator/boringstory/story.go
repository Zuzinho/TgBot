package boringstory

import (
	"ZuzinhoBot/database"
	"ZuzinhoBot/database/selector"
	"strings"
)

func Story() (*strings.Reader, error) {
	story, err := selector.Select(database.BoringStoriesTableName)
	if err != nil {
		return nil, err
	}

	rd := strings.NewReader(story)

	return rd, nil
}
