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
