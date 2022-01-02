package envelopeenc

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

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
	assert.EqualJSON(t, env, `{
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
