package envvar

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func TestParse(t *testing.T) {
	key, val := Parse("FOO=bar")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "bar")
}

func TestParseEmptyValue(t *testing.T) {
	key, val := Parse("FOO=")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "")
}

func TestParseSyntaxError(t *testing.T) {
	key, val := Parse("=")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}

func TestParseEmptyString(t *testing.T) {
	key, val := Parse("")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}
