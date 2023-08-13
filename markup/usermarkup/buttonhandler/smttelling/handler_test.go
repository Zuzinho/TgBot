package smttelling

import "testing"

func TestHandler_Handle(t *testing.T) {
	handler := &Handler{}
	msgSlice, err := handler.Handle(1)
	if err != nil {
		t.Fatal(err)
	}

	for _, msg := range msgSlice {
		t.Logf("Message text - '%s'\n", msg.Text)
	}
}
