package envelopeenc

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestEncryptFailsWithoutSlots(t *testing.T) {
	_, err := Encrypt([]byte("sekrit"), []SlotEncrypter{}, "label")
	assert.Equal(t, err.Error(), "envelope with zero slots not supported")
}
