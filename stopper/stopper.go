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

	// waits for all stoppers to call .Done()
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
	// why SignalReceived? consider this case where one worker stop needs two different places
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
	//         if stop.SignalReceived { // error expected because we ourself .Close()d it
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
	SignalReceived   bool
	Signal           chan bool // each worker must read exactly one ShouldStop signal
	workersWaitGroup *sync.WaitGroup
}

func (s *Stopper) Done() {
	s.workersWaitGroup.Done()
	s.workersWaitGroup = nil // prevent calling Done() twice
}
