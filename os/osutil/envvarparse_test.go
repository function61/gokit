package osutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestParseEnv(t *testing.T) {
	key, val := ParseEnv("FOO=bar")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "bar")
}

func TestParseEnvEmptyValue(t *testing.T) {
	key, val := ParseEnv("FOO=")

	assert.EqualString(t, key, "FOO")
	assert.EqualString(t, val, "")
}

func TestParseEnvSyntaxError(t *testing.T) {
	key, val := ParseEnv("=")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}

func TestParseEnvEmptyString(t *testing.T) {
	key, val := ParseEnv("")

	assert.EqualString(t, key, "")
	assert.EqualString(t, val, "")
}
