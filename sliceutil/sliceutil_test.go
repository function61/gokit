package sliceutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestContainsString(t *testing.T) {
	assert.Equal(t, ContainsString([]string{"foo", "bar"}, "bar"), true)
	assert.Equal(t, ContainsString([]string{"foo", "bar"}, "ufo"), false)
}

func TestFilterString(t *testing.T) {
	result := FilterString([]string{"one", "two", "three", "four", "five"}, func(item string) bool {
		return item != "three" && item != "four"
	})

	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "one")
	assert.Equal(t, result[1], "two")
	assert.Equal(t, result[2], "five")
}

func TestContainsInt(t *testing.T) {
	assert.Equal(t, ContainsInt([]int{1, 2}, 2), true)
	assert.Equal(t, ContainsInt([]int{1, 2}, 3), false)
}

func TestFilterInt(t *testing.T) {
	result := FilterInt([]int{1, 2, 3, 4, 5}, func(item int) bool {
		return item >= 4
	})

	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], 4)
	assert.Equal(t, result[1], 5)
}
