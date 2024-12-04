package osutil

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

// if "err" is non-nil, exits with exit code 1 and prints the given error to stderr
//
// intended to be ran once from program's entrypoint / main function
//
// this abstraction seems silly, but I've written this so many times, and some of the times
// wrong (stdout instead of stderr), that maybe it's better to write this properly once
func ExitIfError(err error) {
	if err == nil { // happy case :)
		return
	}

	errorStream := os.Stderr

	if isatty.IsTerminal(errorStream.Fd()) {
		// non-technical folk more often have trouble recognizing error condition from output, so
		// better make it STAND OUT that the following is an error message
		fmt.Fprintf(errorStream, "✗ ERROR: %s\n", err.Error())
	} else { // little less decoration for machine-to-machine communication
		fmt.Fprintf(errorStream, "ERROR: %s\n", err.Error())
	}

	os.Exit(1)
}
