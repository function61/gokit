package taskrunner

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

func TestWaitThreeTasks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	runner := New(ctx, discardLogger())

	taskStarted := time.Now()

	// the below tasks don't stop immediately

	runner.Start("stopping takes 20ms", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		time.Sleep(20 * time.Millisecond)
		return nil
	})

	runner.Start("stopping takes 40ms", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		time.Sleep(40 * time.Millisecond)
		return nil
	})

	runner.Start("stopping takes 60ms", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		time.Sleep(60 * time.Millisecond)
		return nil
	})

	time.Sleep(10 * time.Millisecond)

	// now expect all tasks to have exited after 70ms mark
	cancel()

	assert.Ok(t, runner.Wait())
	assert.Equal(t, time.Since(taskStarted) > 69*time.Millisecond, true)
}

func TestCancellationStopsTask(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	runner := New(ctx, discardLogger())

	taskStarted := time.Now()

	runner.Start("hangForever", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		return nil
	})

	time.Sleep(60 * time.Millisecond)

	cancel()

	assert.Ok(t, <-runner.Done())

	assert.Equal(t, time.Since(taskStarted) > 59*time.Millisecond, true)
}

func TestStoppingFails(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	runner := New(ctx, discardLogger())

	runner.Start("hangForever", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		return errors.New("i failed stopping :(")
	})

	cancel()

	assert.Equal(t, runner.Wait().Error(), "task 'hangForever' exited with: i failed stopping :(")
}

func TestTaskErrorShouldStopSiblings(t *testing.T) {
	runner := New(context.Background(), discardLogger())

	correctlyWorkingSiblingAlsoStopped := false

	runner.Start("hangForever", func(taskCtx context.Context) error {
		<-taskCtx.Done()
		correctlyWorkingSiblingAlsoStopped = true
		return nil
	})

	runner.Start("fails", func(_ context.Context) error {
		return errors.New("i fail immediately")
	})

	assert.Equal(t, runner.Wait().Error(), "unexpected exit of fails: i fail immediately")
	assert.Equal(t, correctlyWorkingSiblingAlsoStopped, true)
}

func discardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
