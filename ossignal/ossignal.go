package ossignal

import (
	"os"
	"os/signal"
	"syscall"
)

// "The SIGINT signal is sent to a process by its controlling terminal when a user wishes to interrupt the process"
// "The SIGTERM signal is sent to a process to request its termination"
// "SIGINT is nearly identical to SIGTERM"
//     https://en.wikipedia.org/wiki/Signal_(IPC)

func InterruptOrTerminate() <-chan os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	return ch
}
