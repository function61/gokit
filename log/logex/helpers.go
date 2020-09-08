package logex

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// a logger that discards its output
var Discard = log.New(ioutil.Discard, "", 0)

// helper for libraries in getting an alwas non-nil logger. if a logger given to you was
// nil, you get a Discard logger back, which you can call safely - it doesn't write anywhere
func NonNil(ref *log.Logger) *log.Logger {
	if ref == nil {
		return Discard
	}

	return ref
}

// centralized place for creating "standard" stderr logger in a way that supports suppressing
// log timestamps if a mechanism around it (Docker/systemd) would add one anyway
func StandardLogger() *log.Logger {
	return StandardLoggerTo(os.Stderr)
}

// same as StandardLogger() but with explicit sink
func StandardLoggerTo(sink io.Writer) *log.Logger {
	flags := log.LstdFlags

	if os.Getenv("LOGGER_SUPPRESS_TIMESTAMPS") == "1" {
		flags = 0 // LstdFlags were "Ldate | Ltime"
	}

	return log.New(sink, "", flags)
}
