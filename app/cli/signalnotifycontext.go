package cli

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// same as `signal.NotifyContext()` but logs output when got signal to give useful feedback to user
// that we have begun teardown. this feedback becomes extremely important if the teardown process
// takes time or gets stuck.
func notifyContextInterruptOrTerminate() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// need a buffered channel because the sending side is non-blocking
	ch := make(chan os.Signal, 1)

	// "The SIGINT signal is sent to a process by its controlling terminal when a user wishes to interrupt the process"
	// "The SIGTERM signal is sent to a process to request its termination"
	// "SIGINT is nearly identical to SIGTERM"
	//     https://en.wikipedia.org/wiki/Signal_(IPC)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("STOPPING. (If stuck, send sig again to force exit.)", "signal", <-ch)

		// stop accepting signals on the channel. this undoes the effect of this func,
		// and thus makes the process terminate on the next received signal (so you can stop
		// your program if the cleanup code gets stuck)
		signal.Stop(ch)

		cancel()
	}()

	return ctx
}
