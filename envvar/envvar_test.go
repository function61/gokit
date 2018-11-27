package envvar

import (
	"github.com/function61/gokit/assert"
	"os"
	"testing"
)

func TestEnvThatDoesNotExist(t *testing.T) {
	_, err := Get("DOES_NOT_EXIST")

	assert.EqualString(t, err.Error(), "ENV not defined: DOES_NOT_EXIST")
}

func TestEnvThatExists(t *testing.T) {
	val, err := Get("PATH")

	assert.Assert(t, err == nil)
	assert.Assert(t, len(val) > 1)
}

func TestBase64Encoded(t *testing.T) {
	os.Setenv("FOO", "aGkgbW9tIQ==")

	_, err := GetFromBase64Encoded("FOO2")

	assert.EqualString(t, err.Error(), "ENV not defined: FOO2")

	fooVal, err := GetFromBase64Encoded("FOO")

	assert.Assert(t, err == nil)
	assert.EqualString(t, string(fooVal), "hi mom!")
}
