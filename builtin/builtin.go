// Helpers so generic & useful they "could/should be" language-level features.
// This package is intended to be imported with .  so they can be used in non-qualified way.
package builtin

import (
	"context"
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
