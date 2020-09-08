package stringutils

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestSubstr(t *testing.T) {
	assert.EqualString(t, Substr("The quick brown fox jumps over the lazy dog", 0, 3), "The")
	assert.EqualString(t, Substr("The quick brown fox jumps over the lazy dog", 4, 5), "quick")
	assert.EqualString(t, Substr("foobar", 666, 5), "")
	assert.EqualString(t, Substr("foobar", 3, 2), "ba")
	assert.EqualString(t, Substr("foobar", 3, 666), "bar")
	assert.EqualString(t, Substr("foobar", 0, 6), "foobar")
	assert.EqualString(t, Substr("foobar", 0, 666), "foobar")
}

func TestTruncate(t *testing.T) {
	assert.EqualString(t, Truncate("foobar", 6), "foobar")
	assert.EqualString(t, Truncate("foobar", 5), "foo..")
	assert.EqualString(t, Truncate("foobar", 3), "f..")
	assert.EqualString(t, Truncate("foobar", 2), "..")
	assert.EqualString(t, Truncate("foobar", 1), "..")
	assert.EqualString(t, Truncate("foobar", 0), "..")
}
