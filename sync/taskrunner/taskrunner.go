// Supports starting many cancellable goroutines whose exit and their error codes we can wait on.
//
// Works similar to golang.org/x/sync/errgroup but provides richer logging (tasks are
// named) and errors
//
//   - The tasks are expected to run forever, until context cancellation (e.g. task stopping
//     before cancellation even with nil error is considered an error).
//   - If any of the tasks fail, sibling tasks are canceled as well.
package taskrunner

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type Runner struct {
	runningTasks   []*taskItem
	ctx            context.Context // canceled by parent context or by us if any sibling task fails
	cancelAllTasks context.CancelFunc
	taskExited     chan *taskItem
	log            *slog.Logger
	allTasksExited chan error
	startWaiting   sync.Once
}

func New(ctx context.Context, logger *slog.Logger) *Runner {
	ctx, cancel := context.WithCancel(ctx)

	return &Runner{
		runningTasks:   []*taskItem{},
		ctx:            ctx,
		cancelAllTasks: cancel,
		log:            logger,
		taskExited:     make(chan *taskItem),
		allTasksExited: make(chan error),
	}
}

// not safe to be called after you call Wait() or Done()
// not safe for concurrent use
func (t *Runner) Start(taskName string, fn func(ctx context.Context) error) {
	t.log.Debug("starting", "task", taskName)

	task := &taskItem{name: taskName}

	t.runningTasks = append(t.runningTasks, task)

	go func() {
		task.result = fn(t.ctx)

		t.taskExited <- task
	}()
}

// same semantics as Wait(), but returns chan so you can do other stuff while waiting.
//
// NOTE: same note as for Wait()
func (t *Runner) Done() <-chan error {
	t.startWaiting.Do(func() {
		go func() {
			t.allTasksExited <- t.waitInternal()
			close(t.allTasksExited)
		}()
	})

	return t.allTasksExited
}

// Returns:
// - nil if all task exits were expected and errors were nil
// - err if any of the tasks exited unexpectedly or exit was expected but errored
//
// NOTE: don't have multiple goroutines concurrently use Done() or Wait(). i.e. you can
// call Done() and even after that call Wait(), but don't do it concurrently (the
// channel only sends one value)
func (t *Runner) Wait() error {
	return <-t.Done()
}

// the caller has ensured that there is only one instance of this function running
func (t *Runner) waitInternal() error {
	// this is Wait()'s return value as described in method comment
	var allTasksResult error // will be `nil` in the happy case

	processOneExit := func(task *taskItem) {
		t.removeFromRunningTasks(task)

		// if any of the tasks exited with error (even if that was after we were requested to stop),
		// the aggregate result will be an error.
		if err := task.result; err != nil {
			t.log.Error("exited with error", "task", task.name, "err", err)

			if allTasksResult == nil {
				allTasksResult = fmt.Errorf("task '%s' exited with: %w", task.name, err)
			}
		} else {
			t.log.Debug("exited", "task", task.name)
		}
	}

	// ensures that all running tasks have exited and `processOneExit()` has been called on each of them
	waitRunningTasksToExit := func() {
		for len(t.runningTasks) > 0 {
			select {
			case <-time.After(3 * time.Second):
				for _, stuckTask := range t.runningTasks {
					t.log.Warn("waiting for stuck task to exit", "task", stuckTask.name)
				}

				// go back to waiting
			case exit := <-t.taskExited:
				processOneExit(exit)
			}
		}
	}

	select {
	case <-t.ctx.Done():
		waitRunningTasksToExit()
		return allTasksResult
	case maybeUnexpectedExit := <-t.taskExited: // handle unexpected exits
		// why maybeUnexpectedExit?
		// there might be race when this runner's ctx is cancelled because its cancellation
		// is propagated to child contexts (= tasks) so a task can exit and we end up here
		// even though it's not an unexpected exit
		select {
		case <-t.ctx.Done(): // handle the race here
			definitelyExpectedExit := maybeUnexpectedExit
			processOneExit(definitelyExpectedExit)
			waitRunningTasksToExit()
			return allTasksResult
		default:
			// was not a race, continue
		}

		// got unexpected exit => bring sibling tasks down because they all fail as one
		unexpectedExit := maybeUnexpectedExit

		t.removeFromRunningTasks(unexpectedExit)

		t.log.Error("unexpected exit", "err", unexpectedExit.result, "task", unexpectedExit.name)

		err := fmt.Errorf("unexpected exit of %s: %v", unexpectedExit.name, unexpectedExit.result)

		// all sibling tasks fail on first unexpected exit
		t.cancelAllTasks()

		waitRunningTasksToExit()

		return err
	}
}

func (t *Runner) removeFromRunningTasks(task *taskItem) {
	for idx, waiting := range t.runningTasks { // remove from waiting
		if task == waiting {
			t.runningTasks = append(t.runningTasks[:idx], t.runningTasks[idx+1:]...)
			return
		}
	}

	panic(fmt.Errorf("taskrunner removeFromRunningTasks: shouldn't happen: %s", task.name))
}

type taskItem struct {
	name   string
	result error // filled when task exits
}
