package osutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestParseEnv(t *testing.T) {
	key, val := ParseEnv("FOO=bar")

	assert.Equal(t, key, "FOO")
	assert.Equal(t, val, "bar")
}

func TestParseEnvEmptyValue(t *testing.T) {
	key, val := ParseEnv("FOO=")

	assert.Equal(t, key, "FOO")
	assert.Equal(t, val, "")
}

func TestParseEnvSyntaxError(t *testing.T) {
	key, val := ParseEnv("=")

	assert.Equal(t, key, "")
	assert.Equal(t, val, "")
}

func TestParseEnvEmptyString(t *testing.T) {
	key, val := ParseEnv("")

	assert.Equal(t, key, "")
	assert.Equal(t, val, "")
}
