package stopper

import (
	"testing"
	"time"
)

func TestStopper(t *testing.T) {
	stopManager := NewManager()

	// spawn two "workers" who both take 200ms to stop after asking to stop

	go func(stop *Stopper) {
		defer stop.Done()
		<-stop.Signal

		time.Sleep(200 * time.Millisecond)
	}(stopManager.Stopper())

	go func(stop *Stopper) {
		defer stop.Done()
		<-stop.Signal

		time.Sleep(200 * time.Millisecond)
	}(stopManager.Stopper())

	stoppingStarted := time.Now()

	time.Sleep(100 * time.Millisecond)

	stopManager.StopAllWorkersAndWait()

	// above should take 100ms + parallel(200ms, 200ms) = 300ms
	if time.Since(stoppingStarted) < 290*time.Millisecond {
		t.Error("test executed too fast")
	}
}
