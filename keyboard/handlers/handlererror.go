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

type InAccessedHandlerError struct {
	error
	value values.KeyboardValue
}

func NewInAccessedHandlerError(value values.KeyboardValue) InAccessedHandlerError {
	return InAccessedHandlerError{
		value: value,
	}
}

func (err InAccessedHandlerError) String() string {
	return fmt.Sprintf("handler for data '%s' not accessed for not user or admin", err.value)
}
