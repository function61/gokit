package stringutils

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestSubstr(t *testing.T) {
	assert.Equal(t, Substr("The quick brown fox jumps over the lazy dog", 0, 3), "The")
	assert.Equal(t, Substr("The quick brown fox jumps over the lazy dog", 4, 5), "quick")
	assert.Equal(t, Substr("foobar", 666, 5), "")
	assert.Equal(t, Substr("foobar", 3, 2), "ba")
	assert.Equal(t, Substr("foobar", 3, 666), "bar")
	assert.Equal(t, Substr("foobar", 0, 6), "foobar")
	assert.Equal(t, Substr("foobar", 0, 666), "foobar")
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, Truncate("foobar", 6), "foobar")
	assert.Equal(t, Truncate("foobar", 5), "foo..")
	assert.Equal(t, Truncate("foobar", 3), "f..")
	assert.Equal(t, Truncate("foobar", 2), "..")
	assert.Equal(t, Truncate("foobar", 1), "..")
	assert.Equal(t, Truncate("foobar", 0), "..")
}
