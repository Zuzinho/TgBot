package factgenerator

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const src = "http://free-generator.ru/generator.php?action=fact"

type factStruct struct {
	Fact struct {
		Id   string `json:"id"`
		Fact string `json:"fact"`
		Type string `json:"type"`
		Rate string `json:"rate"`
	} `json:"fact"`
}

func Fact() (*strings.Reader, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	f := factStruct{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(f.Fact.Fact), nil
}
