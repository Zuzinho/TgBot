package buttonhandler

import "fmt"

type NoHandlerError struct {
	error
	btnCnt string
}

func NewNoHandlerError(btnCnt string) *NoHandlerError {
	return &NoHandlerError{
		btnCnt: btnCnt,
	}
}

func (err NoHandlerError) String() string {
	return fmt.Sprintf("No handler for button '%s'", err.btnCnt)
}
