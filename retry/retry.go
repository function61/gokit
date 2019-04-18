// Battle-tested algorithm for retrying with timeout and exponential backoff
package retry

import (
	"context"
	"fmt"
	"github.com/function61/gokit/backoff"
	"time"
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

		attemptDuration := time.Since(attemptStarted)

		err := fmt.Errorf("attempt %d failed in %s: %s", attemptNumber, attemptDuration, errAttempt.Error())

		select {
		case <-ctx.Done(): // context canceled? (= deadline exceeded)
			err = fmt.Errorf("GIVING UP (context timeout): %s", err.Error())
			failed(err)
			return err
		case <-time.After(backoffDuration()):
		}

		failed(err)

		attemptNumber++
	}
}

// 0ms, 100 ms, 200 ms, 400 ms, 800 ms, 1000 ms, 1000 ms, ...
func DefaultBackoff() backoff.Func {
	return backoff.ExponentialWithCappedMax(100*time.Millisecond, 1*time.Second)
}
