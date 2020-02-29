// Small HTTP related utils
package httputils

import (
	"context"
	"net/http"
)

func RemoveGracefulServerClosedError(httpServerError error) error {
	if httpServerError == http.ErrServerClosed {
		return nil
	} else {
		// some other error
		// (or nil, but http server should always exit with non-nil error)
		return httpServerError
	}
}

// helper for making HTTP shutdown task (as in compatible with gokit/taskrunner)
//
// Go's http.Server is weird that we cannot use context cancellation to stop it, but instead
// we have to call srv.Shutdown()
func ServerShutdownTask(server *http.Server) func(context.Context, string) error {
	return func(ctx context.Context, _ string) error {
		<-ctx.Done()
		// can't use task ctx b/c it'd cancel the Shutdown() itself
		return server.Shutdown(context.Background())
	}
}
