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
