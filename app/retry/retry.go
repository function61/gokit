// Battle-tested algorithm for retrying with timeout and exponential backoff
package retry

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/function61/gokit/app/backoff"
)

func Retry(
	ctx context.Context,
	attempt func(ctx context.Context) error,
	backoffDuration backoff.Func,
	failed func(err error),
) error {
	attemptNumber := 1

	for {
		attemptStarted := time.Now()

		errAttempt := attempt(ctx)
		if errAttempt == nil {
			return nil // no error, happy path
		}

		errAttemptStructural := &retryAttemptError{
			attemptError:    errAttempt,
			attemptDuration: time.Since(attemptStarted),
			attemptNumber:   attemptNumber,
		}

		failed(errAttemptStructural)

		select {
		case <-ctx.Done(): // context canceled? (= deadline exceeded)
			// TODO: what about explicit cancel (not timeout?)

			// not calling `failed()` here because the surfacing of this conclusion error is the
			// caller's responsibility.

			return &retryAggregateFailed{
				outcome:     fmt.Errorf("Retry: bailing out (%w)", ctx.Err()),
				lastAttempt: errAttemptStructural,
			}
		case <-time.After(backoffDuration()):
		}

		attemptNumber++
	}
}

// 0ms, 100 ms, 200 ms, 400 ms, 800 ms, 1000 ms, 1000 ms, ...
func DefaultBackoff() backoff.Func {
	return backoff.ExponentialWithCappedMax(100*time.Millisecond, 1*time.Second)
}

// creates `Retry()` compatible error handler that only starts reacting to errors after a specific initial duration
func IgnoreErrorsWithin(expectErrorsWithin time.Duration, handleError func(error)) func(error) {
	retryGroupStarted := time.Now()

	return func(err error) {
		// during this time window attempts are expected to fail and thus we should not confuse the user.
		if time.Since(retryGroupStarted) <= expectErrorsWithin {
			return
		}

		// now in this time window we have passed the expected failures and the failures have likely
		// become unexpected. inform the user.

		handleError(err)
	}
}

// creates `Retry()` compatible error handler function that just pushes errors to a logger
func ErrorsToLogger(logger *slog.Logger) func(error) {
	return func(err error) {
		logger.Warn("retry attempt failed", "attempt", err)
	}
}

type retryAggregateFailed struct {
	outcome     error // intentionally does not wrap `lastAttempt`
	lastAttempt error
}

var _ interface {
	error
	slog.LogValuer
} = (*retryAggregateFailed)(nil)

func (r *retryAggregateFailed) Error() string {
	return fmt.Sprintf("%v: %v", r.outcome, r.lastAttempt)
}

func (r *retryAggregateFailed) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("outcome", r.outcome.Error()),
		slog.Any("last_attempt", r.lastAttempt),
	)
}

// this exists mainly to be able to unpack these details in bowels of `ErrorsToLogger()` via `error` type.
// (to be able to structurally log the details.)
type retryAttemptError struct {
	attemptError    error
	attemptDuration time.Duration
	attemptNumber   int
}

var _ interface {
	error
	slog.LogValuer
} = (*retryAttemptError)(nil)

func (r *retryAttemptError) Error() string {
	return fmt.Sprintf("attempt %d failed (in %s): %v", r.attemptNumber, r.attemptDuration, r.attemptError.Error())
}

func (r *retryAttemptError) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("number", r.attemptNumber),
		slog.Any("error", r.attemptError),
		slog.Any("duration", r.attemptDuration),
	)
}
