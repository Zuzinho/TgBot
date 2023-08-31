package selector

import (
	"ZuzinhoBot/database/othertype/statistictype"
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"log"
	"strings"
	"time"
)

func SelectStory(dataType statistictype.DataType) (*strings.Reader, error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var story string
	err = conn.QueryRow("select * from get_story($1)", dataType).Scan(&story)
	if err != nil {
		return nil, err
	}

	log.Printf("Got story with type %s - %s", dataType, story)

	return strings.NewReader(story), nil
}

func SelectUserRole(chatId int64, userName string) (role statistictype.Role, err error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow("select role from users where id = $1 and name = $2", chatId, userName).Scan(&role)

	log.Printf("Got from table 'users' role %s", role)

	return
}

func SelectSortedUsersStatistic(isSince bool) (statistictype.SortedArr, error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("select * from get_sorted_users_statistic($1)", isSince)
	if err != nil {
		return nil, err
	}

	return getArrFromRows(rows)
}

func SelectSortedUsersStatisticByDataType(dataType statistictype.DataType) (statistictype.SortedArr, error) {
	conn, err := pgx.Connect(env.ConnConfig)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("select * from get_sorted_users_statistic_by_data_type($1)", dataType)
	if err != nil {
		return nil, err
	}

	return getArrFromRows(rows)
}

func getArrFromRows(rows *pgx.Rows) (statistictype.SortedArr, error) {
	arr := make(statistictype.SortedArr, 0)
	for rows.Next() {
		var date time.Time
		var userName string
		var userRole statistictype.Role
		var dataType statistictype.DataType
		var count int

		err := rows.Scan(&date, &userName, &userRole, &dataType, &count)
		if err != nil {
			return arr, err
		}

		arr = append(arr, statistictype.NewSortedUsersStatisticRow(date, userName, userRole, dataType, count))
	}

	return arr, nil
}
