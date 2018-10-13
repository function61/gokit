package bidipipe

import (
	"fmt"
	"io"
	"sync"
)

func Pipe(party1 io.ReadWriteCloser, party1Name string, party2 io.ReadWriteCloser, party2Name string) error {
	allIoFinished := &sync.WaitGroup{}

	// we have two goroutines Copy() to different directions. either side could encounter
	// the error either from its read (source) or write (destination) side, so irregardless
	// of direction we should try to close both ends
	closeBothEnds := func() {
		party1.Close()
		party2.Close()
	}

	// take note of the first error to happen when Copy()ing to both directions. since when
	// encountering error, we close both ends and that Close() will cause other (expected) error
	// so we are not interested in that. "first error" can be nil, which means clean EOF
	firstErrorCh := make(chan error, 2)

	go func(done *sync.WaitGroup) {
		defer done.Done()

		_, errCopyToParty1 := io.Copy(party1, party2)

		if errCopyToParty1 != nil {
			firstErrorCh <- fmt.Errorf(
				"bidipipe: %s -> %s error: %s",
				party2Name,
				party1Name,
				errCopyToParty1.Error())
		} else { // clean EOF
			firstErrorCh <- nil
		}

		closeBothEnds()
	}(waiterReference(allIoFinished))

	go func(done *sync.WaitGroup) {
		defer done.Done()

		_, errCopyToParty2 := io.Copy(party2, party1)

		if errCopyToParty2 != nil {
			firstErrorCh <- fmt.Errorf(
				"bidipipe: %s -> %s error: %s",
				party1Name,
				party2Name,
				errCopyToParty2.Error())
		} else {
			firstErrorCh <- nil
		}

		closeBothEnds()
	}(waiterReference(allIoFinished))

	allIoFinished.Wait()

	select {
	case firstError := <-firstErrorCh:
		return firstError
	default:
		return nil // no error - this probably shouldnt happen
	}
}

func waiterReference(wg *sync.WaitGroup) *sync.WaitGroup {
	wg.Add(1)
	return wg
}
