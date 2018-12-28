package throttle

import (
	"github.com/function61/gokit/assert"
	"testing"
	"time"
)

func TestBurstThrottler(t *testing.T) {
	timesCalled := 0

	fn := func() {
		timesCalled++
	}

	interval := 100 * time.Millisecond

	throttler, cancel := BurstThrottler(3, interval)
	defer cancel()

	started := time.Now()

	// six invocations should take exactly one interval
	for i := 0; i < 6; i++ {
		throttler(fn)
	}

	elapsed := time.Since(started)

	assert.Assert(t, timesCalled == 6)
	assert.Assert(t, elapsed > 1*interval)
	assert.Assert(t, elapsed < 2*interval)
}
