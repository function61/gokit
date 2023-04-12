package syncutil

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

func TestConcurrently(t *testing.T) {
	work := make(chan uint64)

	workerIdMax := uint64(0)
	sum := uint64(0)

	err := Concurrently(context.Background(), 3, func(ctx context.Context) error {
		workerId := atomic.AddUint64(&workerIdMax, 1)

		for num := range work {
			time.Sleep(10 * time.Millisecond)

			// these should appear in groups of 3, though you'll have to increase the sleep
			// to notice it
			fmt.Printf("%d(%d)\n", workerId, num)

			atomic.AddUint64(&sum, num)
		}

		return nil
	}, func(workersCancel context.Context) {
		defer close(work)

		for i := uint64(0); i < 10; i++ {
			work <- i
		}
	})

	assert.Ok(t, err)
	assert.Equal(t, sum, 0+1+2+3+4+5+6+7+8+9)
}

func TestConcurrently2(t *testing.T) {
	sum := uint64(0)

	err := Concurrently2(context.Background(), 3, func(_ context.Context, num uint64) error {
		time.Sleep(10 * time.Millisecond)

		atomic.AddUint64(&sum, num)

		return nil
	}, ProducerForSlice([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))

	assert.Ok(t, err)
	assert.Equal(t, sum, 0+1+2+3+4+5+6+7+8+9)
}
