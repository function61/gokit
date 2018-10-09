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
		worker.SignalReceived = true
		worker.Signal <- true
	}

	m.workersWaitGroup.Wait()
}

// creates new Stopper, which is the worker's API for getting a stop signal and reporting finish
func (m *Manager) Stopper() *Stopper {
	m.workersWaitGroup.Add(1)

	stopper := &Stopper{
		SignalReceived:   false,
		Signal:           make(chan bool, 1), // buffered so worker recv waiting not required
		workersWaitGroup: m.workersWaitGroup,
	}

	m.workers = append(m.workers, stopper)

	return stopper
}

type Stopper struct {
	// sometimes shutdown logic is in two different places (one place reads stopper.Signal)
	// via channel, and another place that e.g. reads from a socket. now the shutdown causes
	// a socket read error, and that error handling needs to know if the error was to be
	// expected, it can simply query stopper.SignalReceived instead of trying to determine from
	// the error if the socket was closed on purpose or to communicate with the other code
	// that called socket's close()
	SignalReceived   bool
	Signal           chan bool // each worker must read exactly one ShouldStop signal
	workersWaitGroup *sync.WaitGroup
}

func (s *Stopper) Done() {
	s.workersWaitGroup.Done()
	s.workersWaitGroup = nil // prevent calling Done() twice
}
