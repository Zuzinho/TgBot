package insertor

import (
	"ZuzinhoBot/database/othertype/statistictype"
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"time"
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

func InsertUserMessage(userId int64, dataType statistictype.DataType, date time.Time) (err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("insert into users_statistic (user_id, data_type, date) values ($1, $2, $3)",
		userId, dataType, date.Format("2006-01-02"))

	return
}
