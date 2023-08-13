package phrasedb

import (
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"log"
)

func Phrase() (phrase string, err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	err = conn.QueryRow("select content from phrases where id = $1", 1).Scan(&phrase)
	if err != nil {
		return "", err
	}

	log.Printf("From database got phrase '%s'\n", phrase)

	return
}
