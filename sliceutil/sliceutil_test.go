package sliceutil

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func TestContainsString(t *testing.T) {
	assert.Assert(t, ContainsString([]string{"foo", "bar"}, "bar"))
	assert.Assert(t, !ContainsString([]string{"foo", "bar"}, "ufo"))
}

func TestContainsInt(t *testing.T) {
	assert.Assert(t, ContainsInt([]int{1, 2}, 2))
	assert.Assert(t, !ContainsInt([]int{1, 2}, 3))
}

func TestFilterInt(t *testing.T) {
	result := FilterInt([]int{1, 2, 3, 4, 5}, func(item int) bool {
		return item >= 4
	})

	assert.Assert(t, len(result) == 2)
	assert.Assert(t, result[0] == 4)
	assert.Assert(t, result[1] == 5)
}
