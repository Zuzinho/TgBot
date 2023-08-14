package sleepingstorydb

import (
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"log"
)

func Story() (story string, err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	err = conn.QueryRow("select content from sleeping_stories where id = $1", 1).Scan(&story)
	if err != nil {
		return "", err
	}

	log.Printf("From database got story - %s", story)

	return story, nil
}
