package request

import "strings"

type (
	// Request represents partially parsed income request. Command contains main command like SET, GET
	// Args contains a slice of arguments depending on a command
	Request struct {
		Command string
		Args    []string
	}
)

// NewRequestFromMessage parses income message and transforms it into a request struct
func NewRequestFromMessage(msg string) *Request {
	args := strings.Split(msg, " ")

	return &Request{
		Command: args[0],
		Args:    args[1:],
	}
}
