package envelopeenc

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestEncryptFailsWithoutRecipient(t *testing.T) {
	_, _, err := GenerateAndEncryptDek("label")
	assert.EqualString(t, err.Error(), "envelope with zero recipients not supported")
}
