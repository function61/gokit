package sliceutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestContainsString(t *testing.T) {
	assert.Assert(t, ContainsString([]string{"foo", "bar"}, "bar"))
	assert.Assert(t, !ContainsString([]string{"foo", "bar"}, "ufo"))
}

func TestFilterString(t *testing.T) {
	result := FilterString([]string{"one", "two", "three", "four", "five"}, func(item string) bool {
		return item != "three" && item != "four"
	})

	assert.Assert(t, len(result) == 3)
	assert.Assert(t, result[0] == "one")
	assert.Assert(t, result[1] == "two")
	assert.Assert(t, result[2] == "five")
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
