// OS-related utilities
package osutil

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/function61/gokit/log/logex"
)

// makes (background-derived) context that cancels on SIGINT ("Ctrl + c") or SIGTERM. it
// stops signal listening on first received signal, so that if user sends next signal
// (teardown maybe stuck?), it terminates the process immediately.
//
// TODO: will signal.NotifyContext() cover our needs?
func CancelOnInterruptOrTerminate(logger *log.Logger) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// need a buffered channel because the sending side is non-blocking
	ch := make(chan os.Signal, 1)

	// "The SIGINT signal is sent to a process by its controlling terminal when a user wishes to interrupt the process"
	// "The SIGTERM signal is sent to a process to request its termination"
	// "SIGINT is nearly identical to SIGTERM"
	//     https://en.wikipedia.org/wiki/Signal_(IPC)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logex.Levels(logger).Info.Printf("Got %s: STOPPING. (If stuck, send sig again to force exit.)", <-ch)

		// stop accepting signals on the channel. this undoes the effect of this func,
		// and thus makes the process terminate on the next received signal (so you can stop
		// your program if the cleanup code gets stuck)
		signal.Stop(ch)

		cancel()
	}()

	return ctx
}
