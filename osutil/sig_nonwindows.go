// +build !windows

package osutil

import (
	"os"
)

// Windows-supporting process.Signal(os.Interrupt) abstraction
func Interrupt(process *os.Process) error {
	// defer to the standard interface (which sadly doesn't work for Windows)
	return process.Signal(os.Interrupt)
}
