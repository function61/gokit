// Enables fetching app's uptime by saving start time on initialization
package appuptime

import (
	"time"
)

var started = time.Now()

func Elapsed() time.Duration {
	return time.Since(started)
}
