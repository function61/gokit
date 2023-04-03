// Helpers so generic & useful they "could/should be" language-level features.
//
// The name "builtin" is to align with Go's "builtin = predeclared identifiers", see https://pkg.go.dev/builtin
//
// This package is intended to be imported with .  so they can be used in non-qualified way.
package builtin

import (
	"context"
	"fmt"
)

// empty struct with a semantic name, meaning no information is conveyed. example use cases:
// - https://pkg.go.dev/context?utm_source=godoc#Context.Done
// - visitedItems := map[string]Void{}
type Void struct{}

// ignores error in situations where context was canceled and thus errors are to be expected
// and thus are non-interesting
func IgnoreErrorIfCanceled(ctx context.Context, err error) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		return err
	}
}

// returns a pointer to the given value.
//
// Go is so silly, we need to do this.
// AWS has similar helpers: https://pkg.go.dev/github.com/aws/aws-sdk-go/aws#String
func Pointer[T any](val T) *T {
	return &val
}

// helper for wrapping an error with error prefix.
// if error is nil, nil is returned.
// Deprecated: this doesn't carry its own weight - just use `fmt.Errorf()`
func ErrorWrap(prefix string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", prefix, err)
	} else {
		return nil
	}
}

// errors if a field is unset. it is up to the caller to evaluate if the field is unset
func ErrorIfUnset(isUnset bool, fieldName string) error {
	if isUnset {
		return fmt.Errorf("'%s' is required", fieldName)
	} else {
		return nil
	}
}

// returns the first non-empty value.
func FirstNonEmpty[T comparable](values ...T) T {
	var empty T
	for _, value := range values {
		if value != empty {
			return value
		}
	}
	return empty
}

// for a function that returns two values, you can get the value with `Must(fn())` to convert it to
// the value, panicking if producing the value caused an error
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// returns first error, or nil if no errors
// Deprecated: use `FirstNonEmpty()`
func FirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
