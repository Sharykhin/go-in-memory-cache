package cache

import "fmt"

const (
	UnsupportedTypeCode = iota
	KeyDoesNotExistCode
	CorruptedListCode
	CorruptedDictionaryCode
	EmptyList
	SliceBoundsOutOfRange
	WrongNumberOfArguments
	DictionaryDoesNotExist
)

type (
	// Error represents all error within cache package
	// Apart message it contains code to make it easier
	// to understand an error meaning
	Error struct {
		Code    int
		Message string
	}
)

// Error implements error interface
// So Error struct can be used an error return type
func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d. Error message: %s", e.Code, e.Message)
}

// NewError is an error constructor func
func NewError(code int, msg string) error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
