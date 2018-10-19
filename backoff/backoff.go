package backoff

import (
	"time"
)

type Func func() time.Duration

func ExponentialWithCappedMax(base time.Duration, max time.Duration) Func {
	idx := 0

	return func() time.Duration {
		idx++

		if idx == 1 {
			return 0 * time.Millisecond
		}

		sleepDuration := (1 << uint(idx-2)) * base

		if sleepDuration > max {
			sleepDuration = max
		}

		return sleepDuration
	}
}
