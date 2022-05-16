package osutil

import (
	"fmt"
	"os"
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

	// non-technical folk more often have trouble recognizing error condition from message, so
	// better make it STAND OUT that the following is an error message
	fmt.Fprintf(os.Stderr, "✗ ERROR: %s\n", err.Error())
	os.Exit(1)
}
