package cache

import "fmt"

const (
	UnsupportedTypeCode = iota
	KeyDoesNotExistCode
)

type (
	// Error is a common error struct that is used across the storage
	// with an internal code and message
	Error struct {
		Code    int
		Message string
	}
)

// Error implements error interface so we can return our struct like an error
func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d. Error message: %s", e.Code, e.Message)
}

// NewError is a function constructor that create a new instance that satisfies error interface
func NewError(code int, msg string) error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
