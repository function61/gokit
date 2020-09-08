package logex

import (
	"io"
	"io/ioutil"
	"log"
)

// pipes Logger's output (io.Writer) into another Logger
func OutputToAnotherLog(another *log.Logger) io.Writer {
	if another == nil {
		return ioutil.Discard
	}

	return &anotherLogWriter{another}
}

type anotherLogWriter struct {
	another *log.Logger
}

func (d *anotherLogWriter) Write(msg []byte) (int, error) {
	return len(msg), d.another.Output(2, string(msg))
}

func RedirectGlobalStdLog(to *log.Logger) {
	log.SetOutput(OutputToAnotherLog(to))
	// disable prefixes (like timestamp), as the redirected logger configuration is the
	// single source of truth for log message formatting
	log.SetFlags(0)
}
