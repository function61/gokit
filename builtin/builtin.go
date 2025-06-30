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

// ignores error in situations where context was canceled and thus errors are to be expected / normal.
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

// errors if a field is unset. it is up to the caller to evaluate if the field is unset
func ErrorIfUnset(isUnset bool, fieldName string) error {
	if isUnset {
		return fmt.Errorf("'%s' is required", fieldName)
	} else {
		return nil
	}
}

// for a function that returns two values, you can get the value with `Must(fn())` to convert it to
// the value, panicking if producing the value caused an error
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// append to a slice without needing the "assign to same variable" boilerplate
func Append[T any](slice *[]T, item T) {
	*slice = append(*slice, item)
}
