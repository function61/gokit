package syncutil

import (
	"sync"
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

func TestLockAndUnlock(t *testing.T) {
	mu := &sync.Mutex{}

	eventSequence := make(chan string, 3)

	contenderReadyToLock := make(chan bool)
	contenderGotLock := make(chan bool)

	// as a function b/c the defer-based unlocking
	func() {
		defer LockAndUnlock(mu)()

		eventSequence <- "[main] lock acquired"

		go func() {
			close(contenderReadyToLock)
			mu.Lock()
			eventSequence <- "[contenderThread] got lock"
			close(contenderGotLock)
		}()

		<-contenderReadyToLock

		// make sure that the contender is waiting for the lock - if LockAndUnlock() wouldn't
		// Lock(), this would be enough time for contender to Lock() and alter the sequence of events
		time.Sleep(10 * time.Millisecond)

		eventSequence <- "[main] about to release lock"
	}()

	// make sure contender had time to post its event
	<-contenderGotLock

	// this is the true ordering of events
	assert.Equal(t, <-eventSequence, "[main] lock acquired")
	assert.Equal(t, <-eventSequence, "[main] about to release lock")
	assert.Equal(t, <-eventSequence, "[contenderThread] got lock") // critical that this happens after main releases
}
