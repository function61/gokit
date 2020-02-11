package taskrunner

import (
	"context"
	"errors"
	"github.com/function61/gokit/assert"
	"testing"
	"time"
)

func TestWaitThreeTasks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	runner := New(ctx, nil)

	taskStarted := time.Now()

	// the below tasks don't stop immediately

	runner.Start("stopping takes 20ms", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		time.Sleep(20 * time.Millisecond)
		return nil
	})

	runner.Start("stopping takes 40ms", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		time.Sleep(40 * time.Millisecond)
		return nil
	})

	runner.Start("stopping takes 60ms", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		time.Sleep(60 * time.Millisecond)
		return nil
	})

	time.Sleep(10 * time.Millisecond)

	// now expect all tasks to have exited after 70ms mark
	cancel()

	assert.Assert(t, runner.Wait() == nil)
	assert.Assert(t, time.Since(taskStarted) > 69*time.Millisecond)
}

func TestCancellationStopsTask(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	runner := New(ctx, nil)

	taskStarted := time.Now()

	runner.Start("hangForever", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		return nil
	})

	time.Sleep(60 * time.Millisecond)

	cancel()

	assert.Assert(t, <-runner.Done() == nil)

	assert.Assert(t, time.Since(taskStarted) > 59*time.Millisecond)
}

func TestStoppingFails(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	runner := New(ctx, nil)

	runner.Start("hangForever", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		return errors.New("i failed stopping :(")
	})

	cancel()

	assert.EqualString(t, runner.Wait().Error(), "task 'hangForever' exited with: i failed stopping :(")
}

func TestTaskErrorShouldStopSiblings(t *testing.T) {
	runner := New(context.Background(), nil)

	correctlyWorkingSiblingAlsoStopped := false

	runner.Start("hangForever", func(taskCtx context.Context, _ string) error {
		<-taskCtx.Done()
		correctlyWorkingSiblingAlsoStopped = true
		return nil
	})

	runner.Start("fails", func(_ context.Context, _ string) error {
		return errors.New("i fail immediately")
	})

	assert.EqualString(t, runner.Wait().Error(), "unexpected exit of fails: i fail immediately")
	assert.Assert(t, correctlyWorkingSiblingAlsoStopped)
}
