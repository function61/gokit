package netutil

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"sync"

	. "github.com/function61/gokit/builtin"
	"github.com/function61/gokit/sync/syncutil"
)

// *serve* is called on a separate goroutine
func CancelableServe(ctx context.Context, listener net.Listener, serve func(conn net.Conn)) error {
	allDone := sync.WaitGroup{}
	defer allDone.Wait() // drain all outstanding connections

	acceptFailed := syncutil.Async(func() error {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) { // .Close() called
					return nil
				} else {
					return err
				}
			}

			allDone.Add(1)
			go func() {
				defer allDone.Done()

				serve(conn)
			}()
		}
	})

	select {
	case <-ctx.Done(): // caller wants us to stop
		return listener.Close()
	case err := <-acceptFailed:
		_ = listener.Close() // just to be sure

		return err
	}
}

// wraps use of net.Listener with context cancellation.
// ignores listen error if context was canceled (Accept()-related errors are expected then)
func CancelableListen(
	ctx context.Context,
	listener net.Listener,
	listen func() error,
	logger *slog.Logger,
) error {
	go func() {
		<-ctx.Done()

		if err := listener.Close(); err != nil {
			logger.Error("listener close", "err", err)
		}
	}()

	return IgnoreErrorIfCanceled(ctx, listen()) // actual listening happens here
}
