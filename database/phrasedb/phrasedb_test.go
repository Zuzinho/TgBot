package phrasedb

import "testing"

func TestPhrase(t *testing.T) {
	phrase, err := Phrase()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Phrase - '%s'", phrase)
}
