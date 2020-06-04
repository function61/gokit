package osutil

import (
	"os"

	"golang.org/x/sys/windows"
)

// Go doesn't implement sending process.Signal(os.Interrupt) on Windows (though its runtime
// does on the receiving side translate Windows's CTRL_BREAK_EVENT to SIGINT semantics)
func Interrupt(process *os.Process) error {
	return windows.GenerateConsoleCtrlEvent(windows.CTRL_BREAK_EVENT, uint32(process.Pid))
}
