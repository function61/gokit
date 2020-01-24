// Small wrapper for catching SIGINT and SIGTERM
package ossignal

import (
	"context"
	"github.com/function61/gokit/logex"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// "The SIGINT signal is sent to a process by its controlling terminal when a user wishes to interrupt the process"
// "The SIGTERM signal is sent to a process to request its termination"
// "SIGINT is nearly identical to SIGTERM"
//     https://en.wikipedia.org/wiki/Signal_(IPC)
func InterruptOrTerminate() <-chan os.Signal {
	// need a buffered channel so a nonblocking select will catch this
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	return ch
}

// returns (background-derived) context that cancels on SIGINT or SIGTERM
func InterruptOrTerminateBackgroundCtx(logger *log.Logger) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		logex.Levels(logger).Info.Printf("got %s; stopping", <-InterruptOrTerminate())
		cancel()
	}()

	return ctx
}
