// Exponential backoff
package backoff

import (
	"time"
)

type Func func() time.Duration

func ExponentialWithCappedMax(base time.Duration, max time.Duration) Func {
	idx := -1

	return func() time.Duration {
		idx++

		if idx == 0 {
			return 0 * time.Millisecond
		}

		sleepDuration := (1 << uint(idx-1)) * base

		if sleepDuration > max {
			sleepDuration = max
		}

		return sleepDuration
	}
}
