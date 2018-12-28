package throttle

// alternative implementation: https://github.com/golang/go/wiki/RateLimiting

import (
	"sync"
	"time"
)

// BurstThrottler(3, time.Second) => "only let this fn happen 3 times a second"
func BurstThrottler(count int, duration time.Duration) (func(fn func()), func()) {
	capacity := count

	// when duration has elapsed, release more capacity
	moreCapacityAvailable := time.NewTicker(duration)
	cancel := func() {
		moreCapacityAvailable.Stop()
	}

	mu := sync.Mutex{}

	return func(fn func()) {
		mu.Lock()
		defer mu.Unlock()

		if capacity <= 0 {
			<-moreCapacityAvailable.C
			capacity = count
		}

		capacity--

		fn()
	}, cancel
}
