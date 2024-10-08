package syncutil

import (
	"sync"
)

// Think of this as an infinite number of named bathroom stalls. Each named stall can only
// be occupied by one person.
// When you TryLock():
// a) it won't open if it's already occupied. (because it's locked inside) decide to do something else
// b) it opens and you get in and the stall gets reserved/locked for you. When you get out
// you call the unlock callback you obtained from TryLock() to return the stall for use.
type MutexMap struct {
	// value is chan that Lock() can use to listen for unlock event (close of channel)
	locks   map[string]chan bool
	locksMu sync.Mutex // master lock used whenever "locks" is read/written
}

func NewMutexMap() *MutexMap {
	return &MutexMap{
		locks: map[string]chan bool{},
	}
}

func (m *MutexMap) Lock(key string) func() {
	for {
		unlock, tryAgain := m.tryLockInternal(key)
		if tryAgain != nil {
			// wait for someone to unlock (signalled by close of the chan), so we can try
			// locking again (not guaranteed - someone else might try locking same gate)
			<-tryAgain
			continue
		} else {
			return unlock
		}
	}
}

// returns false if gate already open/reserved
// returns true if gate was opened for you. you have to use the returned func to release it
func (m *MutexMap) TryLock(key string) (func(), bool) {
	unlock, tryAgain := m.tryLockInternal(key)
	if tryAgain != nil {
		return unlock, false
	} else {
		return unlock, true
	}
}

// first return is "unlock" function, which will be nil if tryAgain is non-nil
// second return is "tryAgain" whose close you can wait on to to try locking again
func (m *MutexMap) tryLockInternal(key string) (func(), chan bool) {
	m.locksMu.Lock()
	defer m.locksMu.Unlock()

	if tryAgain, open := m.locks[key]; open {
		return nil, tryAgain
	}

	unlocked := make(chan bool)
	m.locks[key] = unlocked

	return func() {
		m.locksMu.Lock()
		defer m.locksMu.Unlock()

		delete(m.locks, key)
		close(unlocked)
	}, nil
}
