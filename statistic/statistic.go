package statistic

import (
	"ZuzinhoBot/database/othertype/statistictype"
	"ZuzinhoBot/database/selector"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"strings"
)

func Statistic(isSince bool) (*strings.Reader, error) {
	rows, err := selector.SelectSortedUsersStatistic(isSince)
	if errors.Is(err, pgx.ErrNoRows) {
		return strings.NewReader("Статистика отсуствует :(("), nil
	} else if err != nil {
		return nil, err
	}

	return convertStatisticArrToReader(rows)
}

func StatisticByDataType(dataType statistictype.DataType) (*strings.Reader, error) {
	rows, err := selector.SelectSortedUsersStatisticByDataType(dataType)
	if errors.Is(err, pgx.ErrNoRows) {
		return strings.NewReader(fmt.Sprintf("Статистика по типу '%s' отсуствует :((", dataType)), nil
	} else if err != nil {
		return nil, err
	}

	return convertStatisticArrToReader(rows)
}

func convertStatisticArrToReader(rows statistictype.SortedArr) (*strings.Reader, error) {
	builder := strings.Builder{}

	log.Println(rows)

	ln := len(rows)
	dateR := dateRow{date: rows[0].Date}
	for i := 0; i < ln; i++ {
		row := rows[i]

		if dateR.date != row.Date {
			builder.WriteString(dateR.String())
			dateR = dateRow{date: row.Date}
		}

		userR := userRow{userName: row.UserName, userRole: row.UserRole}
		for j := i; j < ln; j++ {
			row = rows[j]

			if row.Date != dateR.date {
				dateR.userRows = append(dateR.userRows, userR)
				userR = userRow{userName: row.UserName, userRole: row.UserRole}
				builder.WriteString(dateR.String())
				dateR = dateRow{date: row.Date}
			}

			log.Printf("Date: %s; User: name - '%s', role - '%s'; Data type: %s, Count: %d", row.Date.Format("02:01:2006"), row.UserName, row.UserRole, row.DataType, row.Count)

			if row.UserName != userR.userName {
				dateR.userRows = append(dateR.userRows, userR)
				userR = userRow{userName: row.UserName, userRole: row.UserRole}
			}

			userR.dataRows = append(userR.dataRows, dataRow{data: row.DataType, count: row.Count})

			if j == ln-1 {
				dateR.userRows = append(dateR.userRows, userR)
				break
			}
			i++
		}

		if i == ln-1 {
			builder.WriteString(dateR.String())
		}
	}

	return strings.NewReader(builder.String()), nil
}
