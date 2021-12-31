package builtin

import (
	"context"
	"errors"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestIgnoreErrorIfCanceled(t *testing.T) {
	anError := errors.New("some error")

	// doesn't touch legit errors context is not done
	assert.Assert(t, IgnoreErrorIfCanceled(context.Background(), anError) == anError)

	ctx, cancel := context.WithCancel(context.Background())

	cancel() // errors are now to be expected so should not be considered an error

	assert.Assert(t, IgnoreErrorIfCanceled(ctx, anError) == nil)
}

func TestPointer(t *testing.T) {
	ptr := Pointer("foo")

	assert.EqualString(t, *ptr, "foo")
}

func TestFirstError(t *testing.T) {
	var errA = errors.New("error A")
	var errB = errors.New("error B")

	assert.Assert(t, FirstError() == nil)
	assert.Assert(t, FirstError(nil) == nil)
	assert.Assert(t, FirstError(errA, nil) == errA)
	assert.Assert(t, FirstError(nil, errA) == errA)
	assert.Assert(t, FirstError(errA, errB) == errA)
	assert.Assert(t, FirstError(errB, errA) == errB)
}
