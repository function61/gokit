// Exponential backoff
package backoff

import (
	"sync"
	"time"
)

type Func func() time.Duration

// given (100ms, 1s) => [0ms, 100ms, 200ms, 400ms, 800ms, 1s, 1s, 1s, ...]
func ExponentialWithCappedMax(base time.Duration, max time.Duration) Func {
	idx := uint(0)
	idxMu := sync.Mutex{}

	return func() time.Duration {
		idxMu.Lock()
		defer idxMu.Unlock()

		var multiplier time.Duration
		if idx == 0 {
			multiplier = 0
		} else {
			multiplier = 1 << (idx - 1)
		}

		sleepDuration := multiplier * base

		if sleepDuration > max {
			sleepDuration = max
		} else {
			// if we'd just shift past max forever we'd overflow long before
			// we get to iteration #64
			idx++
		}

		return sleepDuration
	}
}
