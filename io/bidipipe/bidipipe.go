// Bi-directional pipe (for proxying use cases)
package bidipipe

import (
	"fmt"
	"io"
	"sync"

	"golang.org/x/sync/errgroup"
)

// make two parties talk to each other (usually for proxying use cases). for debugging purposes
// you can optionally name the endpoints (see WithName(), Unnamed() )
func Pipe(
	party1 NamedEndpoint,
	party2 NamedEndpoint,
) error {
	party1.nameIfUnnamed("party1")
	party2.nameIfUnnamed("party2")

	// we have two goroutines Copy() to different directions. either goroutine could encounter
	// the error either from its read (source) or write (destination) side, so irregardless
	// of direction we should try to close both ends
	var closeBothOnce sync.Once
	closeBoth := func() {
		closeBothOnce.Do(func() {
			// since when encountering I/O error, we close both ends and at least one Close()
			// is expected to error, so we just ignore close errors
			party1.endpoint.Close()
			party2.endpoint.Close()
		})
	}

	var pipes errgroup.Group

	pipes.Go(func() error {
		defer closeBoth()

		return pipeOneDir(party1, party2)
	})
	pipes.Go(func() error {
		defer closeBoth()

		return pipeOneDir(party2, party1)
	})

	// returns first non-nil error to be returned from either call to pipeOneDir()
	return pipes.Wait()
}

func pipeOneDir(dst NamedEndpoint, src NamedEndpoint) error {
	if _, err := io.Copy(dst.endpoint, src.endpoint); err != nil {
		return fmt.Errorf(
			"bidipipe: %s -> %s error: %s",
			src.name,
			dst.name,
			err.Error())
	}

	return nil
}

type NamedEndpoint struct {
	name     string
	endpoint io.ReadWriteCloser
}

func (n *NamedEndpoint) nameIfUnnamed(name string) {
	if n.name == "" {
		n.name = name
	}
}

func WithName(name string, endpoint io.ReadWriteCloser) NamedEndpoint {
	return NamedEndpoint{name, endpoint}
}

func Unnamed(endpoint io.ReadWriteCloser) NamedEndpoint {
	return WithName("", endpoint)
}
