package database

import "fmt"

type Table string

const (
	PhrasesTableName         Table = "phrases"
	SleepingStoriesTableName       = "sleeping_stories"
	FunnyStoriesTableName          = "funny_stories"
)

func (t Table) IsValid() bool {
	switch t {
	case PhrasesTableName, SleepingStoriesTableName, FunnyStoriesTableName:
		return true
	default:
		return false
	}
}

type NoTableError struct {
	error
	Table Table
}

func (err NoTableError) String() string {
	return fmt.Sprintf("no table %s in database", err.Table)
}
