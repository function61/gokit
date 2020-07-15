package retry

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/function61/gokit/assert"
)

func TestSucceedsRightAway(t *testing.T) {
	ifFailed := func(err error) {
		t.Error("should not have been called")
	}

	attempts := 0

	err := Retry(context.Background(), func(ctx context.Context) error {
		attempts++

		return nil
	}, DefaultBackoff(), ifFailed)

	assert.Ok(t, err)
	assert.Assert(t, attempts == 1)
}

func TestSucceedsOnThirdTry(t *testing.T) {
	receivedErrors := []error{}

	ifFailed := func(err error) {
		receivedErrors = append(receivedErrors, err)
	}

	attempts := 0

	err := Retry(context.Background(), func(ctx context.Context) error {
		attempts++

		if attempts == 1 {
			return errors.New("fails on first try")
		} else if attempts == 2 {
			return errors.New("fails on second as well")
		}

		return nil
	}, DefaultBackoff(), ifFailed)

	assert.Ok(t, err)
	assert.Assert(t, attempts == 3)
	assert.Assert(t, len(receivedErrors) == 2)
	// use regex to work around variable timing ("failed in 270ns")
	assert.Matches(t, receivedErrors[0].Error(), "attempt 1 failed in .+: fails on first try")
	assert.Matches(t, receivedErrors[1].Error(), "attempt 2 failed in .+: fails on second as well")
}

func TestTakesTooLong(t *testing.T) {
	receivedErrors := []error{}

	ifFailed := func(err error) {
		receivedErrors = append(receivedErrors, err)
	}

	attempts := 0

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := Retry(ctx, func(ctx context.Context) error {
		attempts++

		select {
		case <-ctx.Done():
			return errors.New("encountered timeout")
		case <-time.After(1 * time.Second):
			t.Error("should not happen")
			return nil
		}
	}, DefaultBackoff(), ifFailed)

	assert.Assert(t, attempts == 1)
	assert.Assert(t, len(receivedErrors) == 1)
	assert.Assert(t, regexp.MustCompile(`GIVING UP \(context timeout\): attempt 1 failed in .+: encountered timeout`).MatchString(receivedErrors[0].Error()))
	assert.EqualString(t, err.Error(), receivedErrors[0].Error())
}
