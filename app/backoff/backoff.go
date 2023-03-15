// Exponential backoff
package backoff

import (
	"sync"
	"time"
)

// when an operation is reasonably expected to fail at times and retry attempts should be done,
// this function should be called after each attempt to get the new duration before the next attempt should be tried.
//
// usual implementations are expected to return different duration for each call, like [0ms, 100ms, 200ms, 400ms, ...]
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

// uses a fixed duration for backoff (always returns e.g. 1s, 1s, ...)
func Fixed(dur time.Duration) Func {
	return func() time.Duration { return dur }
}
