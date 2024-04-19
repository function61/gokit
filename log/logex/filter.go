package logex

import (
	"io"
	"log"
	"regexp"
)

// pipes Logger's output (io.Writer) into another Logger, but only if entry matches regex
func Filter(filter *regexp.Regexp, matches *log.Logger) *log.Logger {
	flw := &filterLogWriter{filter: filter}

	logger := log.New(flw, "", 0)
	flw.matches = matches
	return logger
}

type filterLogWriter struct {
	filter  *regexp.Regexp
	matches *log.Logger
}

var _ io.Writer = (*filterLogWriter)(nil)

func (d *filterLogWriter) Write(msg []byte) (int, error) {
	if !d.filter.Match(msg) {
		return len(msg), nil
	}

	return len(msg), d.matches.Output(2, string(msg))
}
