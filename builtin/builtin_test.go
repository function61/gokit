package builtin

import (
	"context"
	"errors"
	"strings"
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

func TestAppend(t *testing.T) {
	favoriteThings := []string{"hamburger"}

	Append(&favoriteThings, "food")
	Append(&favoriteThings, "bar")

	assert.EqualString(t, strings.Join(favoriteThings, ", "), "hamburger, food, bar")
}
