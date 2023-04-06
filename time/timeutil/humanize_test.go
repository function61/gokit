package timeutil

import (
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

func TestHumanizeDuration(t *testing.T) {
	d := func(serialized string) time.Duration {
		dur, err := time.ParseDuration(serialized)
		if err != nil {
			panic(err)
		}
		return dur
	}

	days := func(n float64) time.Duration {
		return 24 * time.Hour * time.Duration(n)
	}

	years := func(n float64) time.Duration {
		return days(n * 365)
	}

	tcs := []struct {
		input  time.Duration
		output string
	}{
		{d("0ms"), "0 milliseconds"},
		{d("1ms"), "1 millisecond"},
		{d("499ms"), "499 milliseconds"},
		{d("500ms"), "1 second"},
		{d("1s"), "1 second"},
		{d("29s"), "29 seconds"},
		{d("30s"), "1 minute"},
		{d("29m"), "29 minutes"},
		{d("36m34.20996749s"), "1 hour"},
		{d("89m"), "1 hour"},
		{d("90m"), "2 hours"},
		{d("12h"), "12 hours"},
		{d("36h"), "1 day"},
		{d("48h"), "2 days"},
		{days(6), "6 days"},
		{days(7), "1 week"},
		{days(13), "1 week"},
		{days(14), "2 weeks"},
		{days(30), "4 weeks"},
		{days(31), "1 month"},
		{days(364), "11 months"},
		{days(365), "1 year"},
		{years(9), "9 years"},
		{years(10.01), "1 decade"},
		{years(20.02), "2 decades"},
	}

	for _, tc := range tcs {
		tc := tc // pin

		t.Run(tc.output, func(t *testing.T) {
			assert.Equal(t, HumanizeDuration(tc.input), tc.output)
		})
	}
}
