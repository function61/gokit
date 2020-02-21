package cryptorandombytes

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func TestHex(t *testing.T) {
	assert.Assert(t, len(Hex(2)) == 4)
	assert.Assert(t, len(Hex(4)) == 8)
	assert.Assert(t, len(Hex(8)) == 16)
}

func TestBase64Url(t *testing.T) {
	// base64 has 33.33 % wasted space
	assert.Assert(t, len(Base64Url(3)) == 4)
	assert.Assert(t, len(Base64Url(6)) == 8)
	assert.Assert(t, len(Base64Url(7)) == 10)
}

func TestBase64UrlWithoutLeadingDash(t *testing.T) {
	assert.Assert(t, len(Base64UrlWithoutLeadingDash(3)) == 4)
}
