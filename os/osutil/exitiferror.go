package osutil

import (
	"fmt"
	"os"
)

// if "err" is non-nil, exits with exit code 1 and prints the given error to stderr
func ExitIfError(err error) {
	// abstraction seems silly, but I've written this so many times, and some of the times
	// wrong (stdout instead of stderr), that maybe it's better to write this properly once
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
