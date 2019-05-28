package logger

import (
	"fmt"
	"os"
)

type (
	// Logger is a general interface that describes necessary method for logging
	Logger interface {
		Errorf(format string, v ...interface{})
		Printf(format string, v ...interface{})
	}

	// TerminalLogger is a concrete implementation of Logger interface.
	// As can be guessed this logger logs messages into stdout, stderr
	TerminalLogger struct {
	}
)

// Printf logs a message into stdout
func (l TerminalLogger) Printf(format string, v ...interface{}) {
	_, _ = os.Stdout.WriteString(fmt.Sprintf(format, v...))
}

// Errorf logs a message in this specific case it will be some sort of error to stderr
func (l TerminalLogger) Errorf(format string, v ...interface{}) {
	_, _ = os.Stderr.WriteString(fmt.Sprintf(format, v...))
}

// NewTerminalLogger returns a new instance of terminal logger
func NewTerminalLogger() *TerminalLogger {
	return &TerminalLogger{}
}
