package insertor

import (
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
)

func InsertUnwantedUser(chatId int64, userName string) (err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("insert into users (id, name, role) values (&1, &2, 1)", chatId, userName)

	return
}
