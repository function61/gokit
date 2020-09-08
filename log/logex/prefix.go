package logex

import (
	"log"
)

// Prefix() creates a new Logger whose output is piped to parent logger with a prefix
func Prefix(prefix string, parent *log.Logger) *log.Logger {
	return log.New(OutputToAnotherLog(parent), prefix+" ", 0)
}
