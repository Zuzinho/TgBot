package phrase

import (
	"io"
	"strings"
)

func Phrase() (io.Reader, error) {
	rd := strings.NewReader("Я тебя люблю")

	return rd, nil
}
