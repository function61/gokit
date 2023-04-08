package throttle

import (
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
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

	assert.Equal(t, timesCalled, 6)
	assert.Equal(t, elapsed > 1*interval, true)
	assert.Equal(t, elapsed < 2*interval, true)
}
