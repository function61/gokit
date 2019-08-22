package backoff

import (
	"fmt"
	"github.com/function61/gokit/assert"
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	backoffDuration := ExponentialWithCappedMax(100*time.Millisecond, 1*time.Second)

	oneSec := 1000 * time.Millisecond

	// why so many test case items? I failed so hard that after 40th iteration we started
	// to converge on producing 0s...
	expectations := []time.Duration{
		0 * time.Millisecond,
		100 * time.Millisecond,
		200 * time.Millisecond,
		400 * time.Millisecond,
		800 * time.Millisecond,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec,
		oneSec, // just to be sure, go past iteration 64
		oneSec,
		oneSec,
		oneSec,
	}

	for idx, expectation := range expectations {
		t.Run(fmt.Sprintf("Iteration %d", idx), func(t *testing.T) {
			assert.Assert(t, backoffDuration() == expectation)
		})
	}
}
