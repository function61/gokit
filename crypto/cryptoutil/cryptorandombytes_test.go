package cryptoutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestRandHex(t *testing.T) {
	assert.Equal(t, len(RandHex(2)), 4)
	assert.Equal(t, len(RandHex(4)), 8)
	assert.Equal(t, len(RandHex(8)), 16)
}

func TestBase64Url(t *testing.T) {
	// base64 has 33.33 % wasted space
	assert.Equal(t, len(RandBase64Url(3)), 4)
	assert.Equal(t, len(RandBase64Url(6)), 8)
	assert.Equal(t, len(RandBase64Url(7)), 10)
}

func TestRandBase64UrlWithoutLeadingDash(t *testing.T) {
	assert.Equal(t, len(RandBase64UrlWithoutLeadingDash(3)), 4)
}
