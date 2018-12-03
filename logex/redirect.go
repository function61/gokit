package logex

import (
	"io"
	"log"
)

// pipes Logger's output (io.Writer) into another Logger
type anotherLogWriter struct {
	another *log.Logger
}

func (d *anotherLogWriter) Write(msg []byte) (int, error) {
	return len(msg), d.another.Output(2, string(msg))
}

func OutputToAnotherLog(another *log.Logger) io.Writer {
	return &anotherLogWriter{another}
}

func RedirectGlobalStdLog(to *log.Logger) {
	log.SetOutput(OutputToAnotherLog(to))
	// disable prefixes (like timestamp), as the redirected logger configuration is the
	// single source of truth for log message formatting
	log.SetFlags(0)
}