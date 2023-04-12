// sync-related utils
package syncutil

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// convenience wrapper around errgroup to run a simple task concurrently, and stop each
// worker if any of the tasks encounter an error during run
func Concurrently(
	ctx context.Context,
	concurrency int,
	task func(ctx context.Context) error,
	workSubmitter func(taskCtx context.Context),
) error {
	// if any of the workers error, taskCtx will be canceled.
	// taskCtx will also be canceled if parent ctx cancels.
	errGroup, taskCtx := errgroup.WithContext(ctx)

	for i := 0; i < concurrency; i++ {
		errGroup.Go(func() error {
			return task(taskCtx)
		})
	}

	// expected to close the channel that is producing work for workers
	workSubmitter(taskCtx)

	return errGroup.Wait()
}

// convenience wrapper around errgroup to run a simple task concurrently, and stop each
// worker if any of the tasks encounter an error during run
func Concurrently2[T any](
	ctx context.Context,
	concurrency int,
	consume func(ctx context.Context, task T) error,
	produce func(taskCtx context.Context, work chan T) error,
) error {
	queue := make(chan T)

	// if any of the workers error, taskCtx will be canceled.
	// taskCtx will also be canceled if parent ctx cancels.
	errGroup, taskCtx := errgroup.WithContext(ctx)

	for i := 0; i < concurrency; i++ {
		errGroup.Go(func() error {
			for task := range queue {
				if err := consume(taskCtx, task); err != nil {
					return fmt.Errorf("consume: %w", err)
				}
			}

			return nil
		})
	}

	closeQueueAndWaitConsumersExit := func() error {
		close(queue)
		return errGroup.Wait()
	}

	if err := produce(taskCtx, queue); err != nil {
		_ = closeQueueAndWaitConsumersExit() // the production error is the most meaningful error
		return fmt.Errorf("produce: %w", err)
	}

	return closeQueueAndWaitConsumersExit()
}

// creates a `Concurrently2()` compatible producer that submits tasks to work queue from a slice
func ProducerForSlice[T any](input []T) func(context.Context, chan T) error {
	return func(ctx context.Context, work chan T) error {
		for _, task := range input {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case work <- task:
				// work item submitted :)
			}
		}

		return nil
	}
}
