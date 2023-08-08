package fact

import (
	"io"
	"log"
	"testing"
)

func TestFact(t *testing.T) {
	for i := 0; i < 5; i++ {
		rd, err := Fact()
		if err != nil {
			log.Fatal(err)
		}

		buf, err := io.ReadAll(rd)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(buf))
	}
}
