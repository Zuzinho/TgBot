package handlers

import (
	"ZuzinhoBot/keyboard/values"
	"fmt"
)

type NoHandlerError struct {
	error
	value values.KeyboardValue
}

func NewNoHandlerError(value values.KeyboardValue) NoHandlerError {
	return NoHandlerError{
		value: value,
	}
}

func (err NoHandlerError) String() string {
	return fmt.Sprintf("No handler for such keyboard value as '%s' or its one is unable", err.value)
}
