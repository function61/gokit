package backoff

import (
	"github.com/function61/gokit/assert"
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	backoffDuration := ExponentialWithCappedMax(100*time.Millisecond, 1*time.Second)

	assert.True(t, backoffDuration() == 0*time.Millisecond)
	assert.True(t, backoffDuration() == 100*time.Millisecond)
	assert.True(t, backoffDuration() == 200*time.Millisecond)
	assert.True(t, backoffDuration() == 400*time.Millisecond)
	assert.True(t, backoffDuration() == 800*time.Millisecond)
	assert.True(t, backoffDuration() == 1000*time.Millisecond)
	assert.True(t, backoffDuration() == 1000*time.Millisecond)
}
