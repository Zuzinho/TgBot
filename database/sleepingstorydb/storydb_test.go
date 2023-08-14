package sleepingstorydb

import "testing"

func TestStory(t *testing.T) {
	story, err := Story()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(story)
}
