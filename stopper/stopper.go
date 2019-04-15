// Provides cancellation (same as Go's "context") and waiting for cancellation completion
// (not provided by Go's "context").
package stopper

import (
	"sync"
)

// stopper manager manages multiple worker's stop signals and the waiting of all of them
// to signal that they're stopped cleanly.
// does about same as https://stackoverflow.com/a/20166250
type Manager struct {
	workers          []*Stopper
	workersWaitGroup *sync.WaitGroup
}

func NewManager() *Manager {
	return &Manager{
		workers:          []*Stopper{},
		workersWaitGroup: &sync.WaitGroup{},
	}
}

func (m *Manager) StopAllWorkersAndWait() {
	for _, worker := range m.workers {
		// does not synchronously stop
		worker.SignalStop()
	}

	// waits for all stoppers to call .Done()
	m.workersWaitGroup.Wait()
}

// creates new Stopper, which is the worker's API for getting a stop signal and reporting finish
func (m *Manager) Stopper() *Stopper {
	m.workersWaitGroup.Add(1)

	stopper := &Stopper{
		Signal:           make(chan interface{}),
		workersWaitGroup: m.workersWaitGroup,
	}

	m.workers = append(m.workers, stopper)

	return stopper
}

// A stopper can stop by itself or be asked from outside to stop
type Stopper struct {
	Signal           chan interface{}
	workersWaitGroup *sync.WaitGroup
	signalStopMu     sync.Mutex
}

// owner of stopper must call this when it's done with its work
func (s *Stopper) Done() {
	s.workersWaitGroup.Done()
	s.workersWaitGroup = nil // prevent calling Done() twice succesfully
}

// can be called by Manager or user directly.
// can safely be called multiple times.
func (s *Stopper) SignalStop() {
	s.signalStopMu.Lock()
	defer s.signalStopMu.Unlock()

	if !s.WasSignalled() { // prevent double-close
		close(s.Signal)
	}
}

// why this? consider this case where one worker stop needs two different places
// in code to handle stopping:
//
// conn := getConnection()
//
// go func() {
//     <-stop.Signal
//     conn.Close()
// }()
//
// for { // reads indefinitely from socket
//     _, err := conn.Read(buffer)
//     if err != nil {
//         if stop.WasSignalled() { // error expected because we ourself .Close()d it
//             return // clean shutdown
//         }
//
//         panic(err) // handle unexpected error
//     }
//
// }
//
// if we didn't have stop.SignalReceived, we'd have to have shared variable "cleanShutdown"
// which we'd set before calling Close() and use that variable in Read() error handling
func (s *Stopper) WasSignalled() bool {
	select {
	case <-s.Signal:
		return true
	default:
		return false
	}
}
