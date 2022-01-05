// Helpers so generic & useful they "could/should be" language-level features.
// This package is intended to be imported with .  so they can be used in non-qualified way.
package builtin

import (
	"context"
	"fmt"
)

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

// returns a pointer to the given string.
//
// will soon benefit from generics love.
//
// Go is so silly, we need to do this.
// AWS has similar helpers: https://pkg.go.dev/github.com/aws/aws-sdk-go/aws#String
func Pointer(str string) *string {
	return &str
}

// returns first error, or nil if no errors
func FirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

// errors if a field is unset. it is up to the caller to evaluate if the field is unset
func ErrorIfUnset(isUnset bool, fieldName string) error {
	if isUnset {
		return fmt.Errorf("'%s' is required", fieldName)
	} else {
		return nil
	}
}