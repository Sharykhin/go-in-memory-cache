package errors

import (
	"errors"
	"log"
)

// DeferredFunc is a semantical definition of deferred func that would be used
type DeferredFunc = func() error

// CheckDeferred just checks whether deferred func returns an error and is so just logs it
func CheckDeferred(fn DeferredFunc) {
	err := fn()
	if err != nil {
		log.Printf("deferred call returned an error: %v", err)
	}
}

// New is a alias to a general error construction func
func New(err string) error {
	return errors.New(err)
}
