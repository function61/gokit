// Supports starting many cancellable goroutines whose exit and their error codes we can wait on.
//
// - The tasks are expected to run forever, until context cancellation (e.g. task stopping
//   before cancellation even with nil error is considered an error).
// - If any of the tasks fail, sibling tasks are canceled as well.
package taskrunner

import (
	"context"
	"fmt"
	"github.com/function61/gokit/logex"
	"log"
	"sync"
)

type Runner struct {
	numTasks       int             // incremented per each Start()
	ctx            context.Context // canceled by parent context or by us if any sibling task fails
	cancelAllTasks context.CancelFunc
	taskExited     chan taskExit
	logl           *logex.Leveled
	allTasksExited chan error
	startWaiting   sync.Once
}

func New(ctx context.Context, logger *log.Logger) *Runner {
	ctx, cancel := context.WithCancel(ctx)

	return &Runner{
		numTasks:       0,
		ctx:            ctx,
		cancelAllTasks: cancel,
		logl:           logex.Levels(logger),
		taskExited:     make(chan taskExit),
		allTasksExited: make(chan error),
	}
}

// not safe to be called after you call Wait() or Done()
// not safe for concurrent use
func (t *Runner) Start(taskName string, fn func(ctx context.Context, taskName string) error) {
	t.numTasks++

	t.logl.Debug.Printf("starting %s", taskName)

	go func() {
		err := fn(t.ctx, taskName)

		t.taskExited <- taskExit{taskName, err}
	}()
}

// same semantics as Wait(), but returns chan so you can do other stuff while waiting.
// this is safe to be called more than once, but if you call this you must not call Wait().
func (t *Runner) Done() <-chan error {
	t.startWaiting.Do(func() {
		go func() {
			t.allTasksExited <- t.Wait()
		}()
	})

	return t.allTasksExited
}

// - Call Wait() only once
// - returns nil if all task exits were expected and errors were nil
// - returns err if any of the tasks exited unexpectedly or exit was expected but errored
func (t *Runner) Wait() error {
	// this is Wait()'s return value as described in method comment
	var firstErr error

	processOneExit := func(taskExit taskExit) {
		if err := taskExit.err; err != nil {
			errWrapped := fmt.Errorf("task '%s' exited with: %w", taskExit.taskName, err)
			t.logl.Error.Println(errWrapped.Error())

			if firstErr == nil {
				firstErr = errWrapped
			}
		} else {
			t.logl.Debug.Printf("%s exited", taskExit.taskName)
		}
	}

	drainExits := func(numTasks int) error {
		for i := 0; i < numTasks; i++ {
			processOneExit(<-t.taskExited)
		}

		return firstErr
	}

	select {
	case <-t.ctx.Done():
		return drainExits(t.numTasks)
	case maybeUnexpectedExit := <-t.taskExited: // handle unexpected exits
		// why maybeUnexpectedExit?
		// there might be race when this runner's ctx is cancelled because its cancellation
		// is propagated to child contexts (= tasks) so a task can exit and we end up here
		// even though it's not an unexpected exit
		select {
		case <-t.ctx.Done(): // handle the race here
			definitelyExpectedExit := maybeUnexpectedExit
			processOneExit(definitelyExpectedExit)
			return drainExits(t.numTasks - 1) // -1 because we already consumed one
		default:
			// was not a race, continue
		}

		// got unexpected exit => bring sibling tasks down because they all fail as one
		unexpectedExit := maybeUnexpectedExit

		firstErr = fmt.Errorf("unexpected exit of %s: %v", unexpectedExit.taskName, unexpectedExit.err)

		t.logl.Error.Println(firstErr.Error())

		// all sibling tasks fail on first unexpected exit
		t.cancelAllTasks()

		return drainExits(t.numTasks - 1)
	}
}

type taskExit struct {
	taskName string
	err      error
}
