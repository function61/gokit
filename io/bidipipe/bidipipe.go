// Bi-directional pipe (for proxying use cases)
package bidipipe

import (
	"fmt"
	"io"
	"sync"
)

func Pipe(party1 io.ReadWriteCloser, party1Name string, party2 io.ReadWriteCloser, party2Name string) error {
	allIoFinished := &sync.WaitGroup{}

	// take note of the first error to happen when Copy()ing to both directions. since when
	// encountering error, we close both ends and that Close() will cause other (expected) error
	// so we are not interested in that.
	firstErrorCh := make(chan error, 2)

	go pipeOneDir(party1, party1Name, party2, party2Name, waiterAdd(allIoFinished), firstErrorCh)
	go pipeOneDir(party2, party2Name, party1, party1Name, waiterAdd(allIoFinished), firstErrorCh)

	allIoFinished.Wait()

	select {
	case firstError := <-firstErrorCh:
		return firstError
	default:
		return nil
	}
}

func pipeOneDir(dst io.ReadWriteCloser, dstName string, src io.ReadWriteCloser, srcName string, done *sync.WaitGroup, firstErrorCh chan error) {
	defer done.Done()

	if _, err := io.Copy(dst, src); err != nil {
		firstErrorCh <- fmt.Errorf(
			"bidipipe: %s -> %s error: %s",
			srcName,
			dstName,
			err.Error())
	}

	// we have two goroutines Copy() to different directions. either goroutine could encounter
	// the error either from its read (source) or write (destination) side, so irregardless
	// of direction we should try to close both ends
	src.Close()
	dst.Close()
}

func waiterAdd(wg *sync.WaitGroup) *sync.WaitGroup {
	wg.Add(1)
	return wg
}
