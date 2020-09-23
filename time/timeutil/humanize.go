// Time-related utilities
package timeutil

import (
	"math"
	"strconv"
	"time"
)

func HumanizeDuration(dur time.Duration) string {
	milliseconds := float64(dur.Milliseconds())
	seconds := int(math.Round(milliseconds / (1.0 * 1000.0)))
	minutes := int(math.Round(milliseconds / (60.0 * 1000.0)))
	hours := int(math.Round(milliseconds / (3600.0 * 1000.0)))
	days := int(milliseconds / (86400.0 * 1000.0))
	weeks := int(float64(days) / 7.0)
	months := int(float64(days) / (365 / 12.0))
	years := int(float64(days) / 365.0)
	decades := int(float64(days) / 3652.0)

	maybePluralize := func(num int, singular string) string {
		if num == 1 {
			return strconv.Itoa(num) + " " + singular
		} else {
			// works only for english
			plural := singular + "s"
			return strconv.Itoa(num) + " " + plural
		}
	}

	switch {
	case decades > 0:
		return maybePluralize(decades, "decade")
	case years > 0:
		return maybePluralize(years, "year")
	case months > 0:
		return maybePluralize(months, "month")
	case weeks > 0:
		return maybePluralize(weeks, "week")
	case days > 0:
		return maybePluralize(days, "day")
	case hours > 0:
		return maybePluralize(hours, "hour")
	case minutes > 0:
		return maybePluralize(minutes, "minute")
	case seconds > 0:
		return maybePluralize(seconds, "second")
	default:
		return maybePluralize(int(milliseconds), "millisecond")
	}
}
