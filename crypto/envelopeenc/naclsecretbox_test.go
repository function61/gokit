package envelopeenc

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestNaclSecretBoxEncrypter(t *testing.T) {
	pwdEnc := NaclSecretBoxEncrypter(
		[32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		"bestpassword")

	assert.EqualString(t, pwdEnc.KekId(), "bestpassword")

	env, err := Encrypt([]byte("my secret message"), []SlotEncrypter{pwdEnc}, "example envelope")
	assert.Ok(t, err)

	decrypted, err := env.Decrypt(pwdEnc)
	assert.Ok(t, err)

	assert.EqualString(t, string(decrypted), "my secret message")

	// label must be non-malleable
	env.Label = "malicious label"

	_, err = env.Decrypt(pwdEnc)
	assert.EqualString(t, err.Error(), "secretbox.Open failed")
}
