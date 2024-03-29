package osutil

import (
	"os"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestEnvThatDoesNotExist(t *testing.T) {
	_, err := GetenvRequired("DOES_NOT_EXIST")

	assert.Equal(t, err.Error(), "ENV not defined: DOES_NOT_EXIST")
}

func TestEnvThatExists(t *testing.T) {
	val, err := GetenvRequired("PATH")

	assert.Ok(t, err)
	assert.Equal(t, len(val) > 1, true)
}

func TestGetenvRequiredFromBase64(t *testing.T) {
	os.Setenv("FOO", "aGkgbW9tIQ==")

	_, err := GetenvRequiredFromBase64("FOO2")

	assert.Equal(t, err.Error(), "ENV not defined: FOO2")

	fooVal, err := GetenvRequiredFromBase64("FOO")

	assert.Ok(t, err)
	assert.Equal(t, string(fooVal), "hi mom!")
}
