package database

import "fmt"

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

type NotAccessedTableError struct {
	error
	table TableName
}

func (err NotAccessedTableError) String() string {
	return fmt.Sprintf("table %s is not accessed for not admin", err.table)
}

func NewNotAccessedTableError(table TableName) NotAccessedTableError {
	return NotAccessedTableError{
		table: table,
	}
}
