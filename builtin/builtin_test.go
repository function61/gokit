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

func TestPointerToString(t *testing.T) {
	ptr := Pointer("foo")

	assert.EqualString(t, *ptr, "foo")
}

func TestPointerToInt(t *testing.T) {
	var num int = 314

	assert.Assert(t, *Pointer(num) == 314)
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

func TestErrorWrap(t *testing.T) {
	assert.EqualString(t, ErrorWrap("myprefix", errors.New("unknown error occurred")).Error(), "myprefix: unknown error occurred")
	assert.Assert(t, ErrorWrap("myprefix", nil) == nil)
}

func TestErrorIfUnset(t *testing.T) {
	assert.Assert(t, ErrorIfUnset(false, "foo") == nil)
	assert.EqualString(t, ErrorIfUnset(true, "foo").Error(), "'foo' is required")
}

func TestFirstNonEmptyWithString(t *testing.T) {
	assert.EqualString(t, FirstNonEmpty("foo", "bar"), "foo")
	assert.EqualString(t, FirstNonEmpty("", "bar"), "bar")
	assert.EqualString(t, FirstNonEmpty(""), "")
}

func TestFirstNonEmptyWithInt(t *testing.T) {
	var zero int
	pi := int(314)

	assert.Assert(t, FirstNonEmpty(zero, pi) == 314)
	assert.Assert(t, FirstNonEmpty(pi, zero) == 314)
}

func TestFirstNonEmptyWithError(t *testing.T) {
	var nilError error = nil
	actualError := errors.New("actual")

	assert.EqualString(t, FirstNonEmpty(nilError, actualError).Error(), "actual")
	assert.Assert(t, FirstNonEmpty(nilError, nilError) == nil)
}
