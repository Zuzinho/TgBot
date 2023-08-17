package database

import "fmt"

type TableName string

const (
	PhrasesTableName       TableName = "phrases"
	BoringStoriesTableName TableName = "boring_stories"
	FunnyStoriesTableName  TableName = "funny_stories"
)

func (t TableName) IsValid() bool {
	switch t {
	case PhrasesTableName, BoringStoriesTableName, FunnyStoriesTableName:
		return true
	default:
		return false
	}
}

type NoTableError struct {
	error
	table TableName
}

func (err NoTableError) String() string {
	return fmt.Sprintf("no table %s in database", err.table)
}

func NewNoTableError(table TableName) NoTableError {
	return NoTableError{
		table: table,
	}
}
