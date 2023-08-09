package token

import (
	"flag"
	"log"
)

func MustToken() string {
	token := flag.String(
		"tg_bot_token",
		"",
		"token for access to tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("No token")
	}

	return *token
}
