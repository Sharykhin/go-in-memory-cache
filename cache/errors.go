package cache

import "fmt"

const (
	UnsupportedTypeCode = iota
	KeyDoesNotExistCode
)

type (
	Error struct {
		Code    int
		Message string
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d. Error message: %s", e.Code, e.Message)
}

func NewError(code int, msg string) error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
