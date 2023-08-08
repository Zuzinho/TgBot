package funnystory

import (
	"io"
	"strings"
)

func Story() (io.Reader, error) {
	rd := strings.NewReader("Однажды она\nПерепутала право и лево")

	return rd, nil
}
