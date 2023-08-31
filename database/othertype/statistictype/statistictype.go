package statistictype

import "time"

type DataType string

type Role string

type SortedUsersStatisticRow struct {
	Date     time.Time
	UserName string
	UserRole Role
	DataType DataType
	Count    int
}

func NewSortedUsersStatisticRow(date time.Time, userName string, userRole Role, dataType DataType, count int) *SortedUsersStatisticRow {
	return &SortedUsersStatisticRow{
		Date:     date,
		UserName: userName,
		UserRole: userRole,
		DataType: dataType,
		Count:    count,
	}
}

type SortedArr []*SortedUsersStatisticRow
