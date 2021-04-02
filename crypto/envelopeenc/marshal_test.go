package envelopeenc

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestMarshal(t *testing.T) {
	defer makeCryptoRandReturn(oneKOfZero())()

	recip, err := Symmetric(symmetricTestKey, "test")
	assert.Ok(t, err)

	out, _, err := Marshal("", recip)
	assert.Ok(t, err)

	assert.EqualString(t, string(out), `{"recip":{"chacha20poly1305:test":"AAAAAAAAAAAAAAAAFO4sz1ePijb5fnlQRXrNtUnXj8qQi4jQgQL0XDdlgm4WNyM/49C612nFZ3cKFCiv"}}
GyMWmRdclEjE5MsE06RgNPyc0nQdn69Z5liYUWtUaTo`)

	un, err := Unmarshal(out)
	assert.Ok(t, err)

	_, err = un.Decrypt(recip)
	assert.Ok(t, err)
}

/*
func TestMarshalUnmarshal(t *testing.T) {
	// TODO: use Encrypt() to build this
	out, err := (&Envelope{
		KeySlots: []KeySlot{
			{
				KekId:        "foo",
				DekEncrypted: []byte{0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00},
			},
		},
		EncryptedContent: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
	}).MarshalDONTUSE()
	assert.Ok(t, err)

	env, err := UnmarshalDONTUSE(out)
	assert.Ok(t, err)

	// TODO: fix kind=0, missing label support
	assert.EqualJson(t, env, `{
  "key_slots": [
    {
      "kind": 0,
      "kek_id": "foo",
      "dek_encrypted": "BwYFBAMCAQA="
    }
  ],
  "content": "AAECAwQFBgc="
}`)
}
*/
