package logex

import (
	"io"
	"log"
	"os"
)

// a logger that discards its output
var Discard = log.New(io.Discard, "", 0)

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

	// "This permits invoked processes to safely detect whether their standard output or standard
	// error output are connected to the journal."
	// https://www.freedesktop.org/software/systemd/man/systemd.exec.html#%24JOURNAL_STREAM
	systemdJournal := os.Getenv("JOURNAL_STREAM") != ""

	// explicitly asked, e.g. set by orchestrator when running in Docker with log redirection taken care of
	explicitSuppress := os.Getenv("LOGGER_SUPPRESS_TIMESTAMPS") == "1"

	if systemdJournal || explicitSuppress {
		flags = 0 // LstdFlags were "Ldate | Ltime"
	}

	return log.New(sink, "", flags)
}
