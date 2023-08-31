package statistic

import (
	"ZuzinhoBot/database/othertype/statistictype"
	"fmt"
	"strings"
	"time"
)

type dateRow struct {
	date     time.Time
	userRows []userRow
}

type userRow struct {
	userName string
	userRole statistictype.Role
	dataRows []dataRow
}

type dataRow struct {
	data  statistictype.DataType
	count int
}

func (dataR dataRow) String() string {
	return fmt.Sprintf("________Тип запроса '%s' - %d раз\n", dataR.data, dataR.count)
}

func (userR userRow) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("____Пользователь '%s' с ролью '%s' вызвал:\n", userR.userName, userR.userRole))

	for _, dataR := range userR.dataRows {
		builder.WriteString(dataR.String())
	}

	return builder.String()
}

func (dateR dateRow) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s бота использовали:\n", dateR.date.Format("02:01:2006")))

	for _, userR := range dateR.userRows {
		builder.WriteString(userR.String())
	}

	return builder.String()
}
