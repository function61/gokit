// Small HTTP related utils
package httputils

import (
	"context"
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func NoCacheHeaders(w http.ResponseWriter) {
	// https://stackoverflow.com/a/2068407
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
}

// helper for adapting context cancellation to shutdown the HTTP listener
func CancelableServer(ctx context.Context, srv *http.Server, listener func() error) error {
	shutdownerCtx, cancel := context.WithCancel(ctx)

	shutdownResult := make(chan error, 1)

	// this is the actual shutdowner
	go func() {
		// triggered by parent cancellation
		// (or below for cleanup if ListenAndServe() failed by itself)
		<-shutdownerCtx.Done()

		// can't use parent ctx b/c it'd cancel the Shutdown() itself
		shutdownResult <- srv.Shutdown(context.Background())
	}()

	err := listener()

	// ask shutdowner to stop. this is useful only for cleanup where listener failed before
	// it was requested to shut down b/c parent cancellation didn't happen and thus the
	// shutdowner would still wait.
	cancel()

	if err == http.ErrServerClosed { // expected for graceful shutdown (not actually error)
		return <-shutdownResult // should be nil, unless shutdown fails
	} else {
		// some other error
		// (or nil, but http server should always exit with non-nil error)
		return err
	}
}

// helper for setting JSON header and JSON-marshaling a struct into the HTTP response
func RespondJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// writing this error probably fails, because the probability of above Encode() failing
		// due to broken conn is much more than JSON marshalling failing
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
