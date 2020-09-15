package osutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestParseEnvs(t *testing.T) {
	key, val := ParseEnvs("FOO=bar")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "bar")
}

func TestParseEnvsEmptyValue(t *testing.T) {
	key, val := ParseEnvs("FOO=")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "")
}

func TestParseEnvsSyntaxError(t *testing.T) {
	key, val := ParseEnvs("=")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}

func TestParseEnvsEmptyString(t *testing.T) {
	key, val := ParseEnvs("")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}
