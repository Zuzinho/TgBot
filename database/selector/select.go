package selector

import (
	"ZuzinhoBot/database"
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"log"
)

func SelectStory(table database.TableName) (value string, err error) {
	if !table.IsValid() {
		return "", database.NewNoTableError(table)
	}

	if !table.IsAccessedForUser() {
		return "", database.NewNotAccessedTableError(table)
	}

	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var id int

	err = conn.QueryRow("select id, content from "+string(table)+" where is_read = false order by RANDOM() limit 1").Scan(&id, &value)
	if err == pgx.ErrNoRows {
		log.Printf("Read all story from table %s", table)

		err = conn.QueryRow("select content from " + string(table) + " order by RANDOM() limit 1").Scan(&value)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		_, err = conn.Exec("update "+string(table)+" set is_read = true where id = $1", id)
		if err != nil {
			return "", err
		}
	}

	log.Printf("From table '%s' got story - %s", table, value)

	return value, nil
}

func SelectUserRole(chatId int64, userName string) (role int, err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow("select role from users where id = $1 and name = $2", chatId, userName).Scan(&role)

	log.Printf("Got from table 'users' role %d", role)

	return
}
