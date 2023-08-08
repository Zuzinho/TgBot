package sleepingstory

import (
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	rd := strings.NewReader("Not\nInteresting\nStory")

	return rd, nil
}
