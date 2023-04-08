package syncutil

import (
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

func TestMutexMap(t *testing.T) {
	mm := NewMutexMap()

	unlockFoo := mm.Lock("foo")
	unlockFoo()

	unlockFoo, fooOk := mm.TryLock("foo")
	assert.Equal(t, fooOk, true)

	_, fooConcurrentOk := mm.TryLock("foo")
	assert.Equal(t, fooConcurrentOk, false)

	unlockFoo()

	unlockFoo, fooOk = mm.TryLock("foo")
	assert.Equal(t, fooOk, true)

	lockAcquireDuration := make(chan time.Duration)

	go func() {
		startedAcquiringLock := time.Now()

		unlock := mm.Lock("foo")
		defer unlock()

		lockAcquireDuration <- time.Since(startedAcquiringLock)
	}()

	time.Sleep(11 * time.Millisecond)

	unlockFoo()

	assert.Equal(t, <-lockAcquireDuration > 10*time.Millisecond, true)
}
