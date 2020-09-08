package envvar

import (
	"os"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestEnvThatDoesNotExist(t *testing.T) {
	_, err := Required("DOES_NOT_EXIST")

	assert.EqualString(t, err.Error(), "ENV not defined: DOES_NOT_EXIST")
}

func TestEnvThatExists(t *testing.T) {
	val, err := Required("PATH")

	assert.Ok(t, err)
	assert.Assert(t, len(val) > 1)
}

func TestBase64Encoded(t *testing.T) {
	os.Setenv("FOO", "aGkgbW9tIQ==")

	_, err := RequiredFromBase64Encoded("FOO2")

	assert.EqualString(t, err.Error(), "ENV not defined: FOO2")

	fooVal, err := RequiredFromBase64Encoded("FOO")

	assert.Ok(t, err)
	assert.EqualString(t, string(fooVal), "hi mom!")
}
