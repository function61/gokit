package syncutil

import (
	"sync"
)

// helper for reducing lock boilerplate by using "defer LockAndUnlock(mu)()" to queue its
// unlocking
func LockAndUnlock(mu *sync.Mutex) func() {
	mu.Lock()

	return func() {
		mu.Unlock()
	}
}
