package stopper

import (
	"sync"
)

// stopper manager manages multiple worker's stop signals and the waiting of all of them
// to signal that they're stopped cleanly.
// does about same as https://stackoverflow.com/a/20166250
type Manager struct {
	workerStopSignal chan bool
	workers          int
	workersWaitGroup *sync.WaitGroup
}

func NewManager() *Manager {
	return &Manager{
		workerStopSignal: make(chan bool, 16),
		workers:          0,
		workersWaitGroup: &sync.WaitGroup{},
	}
}

func (s *Manager) StopAllWorkersAndWait() {
	// fire as many stop messages on the stop signal channel as there are stoppers (workers)
	for i := 0; i < s.workers; i++ {
		s.workerStopSignal <- true
	}

	s.workersWaitGroup.Wait()
}

// creates new Stopper, which is the worker's API for getting a stop signal and reporting finish
func (s *Manager) Stopper() *Stopper {
	s.workers++
	s.workersWaitGroup.Add(1)

	return &Stopper{
		ShouldStop:       s.workerStopSignal,
		workersWaitGroup: s.workersWaitGroup,
	}
}

type Stopper struct {
	ShouldStop       chan bool // each worker must read exactly one ShouldStop signal
	workersWaitGroup *sync.WaitGroup
}

func (s *Stopper) Done() {
	s.workersWaitGroup.Done()
}
