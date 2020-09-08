// Enables fetching app's uptime by saving start time on initialization
package appuptime

import (
	"time"
)

var started = time.Now()

func Started() time.Time {
	return started
}
