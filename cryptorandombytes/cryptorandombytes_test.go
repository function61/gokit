package cryptorandombytes

import (
	"github.com/function61/gokit/assert"
	"testing"
)

func TestHext(t *testing.T) {
	assert.True(t, len(Hex(2)) == 4)
	assert.True(t, len(Hex(4)) == 8)
	assert.True(t, len(Hex(8)) == 16)
}
