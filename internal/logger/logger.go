package logger

import (
	"log"
	"os"
)

const format = log.Ldate | log.Ltime | log.Lshortfile | log.LUTC

// Log the logger for the url shortener service. A basic log package logger with
// an additional function and setting for whether to emit debug information.
type Log struct {
	*log.Logger
	trace bool
}

// MakeLogger creates and returns a logger object appropriate for
// use with the URL shortening service.
// Returns a pointer to a new logger object.
func MakeLogger(trace bool) *Log {
	internalLogger := log.New(os.Stdout, "", format)
	return &Log{
		internalLogger,
		trace,
	}
}

// InfoTracef supports optional logging of formatted non-error type messages. The method will
// not log messages if trace logging was not enabled as application startup.
func (log *Log) InfoTracef(format string, v ...interface{}) {
	if log.trace {
		log.Printf(format, v)
	}
}

// InfoTrace supports optional logging of non-error type messages. The method will
// not log messages if trace logging was not enabled as application startup.
func (log *Log) InfoTrace(v ...interface{}) {
	if log.trace {
		log.Print(v)
	}
}
