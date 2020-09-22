package envelopeenc

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

type naclSecretBox struct {
	kek   [32]byte
	kekId string
}

// symmetric, so this same interface works in both directions
var _ SlotEncrypter = (*naclSecretBox)(nil)
var _ SlotDecrypter = (*naclSecretBox)(nil)

// grant access to envelope if user knows a symmetric 256-bit key (identified by KEK ID)
func NaclSecretBoxEncrypter(kek [32]byte, kekId string) *naclSecretBox {
	return &naclSecretBox{
		kek:   kek,
		kekId: kekId,
	}
}

func (r *naclSecretBox) KekId() string {
	return r.kekId
}

func (r *naclSecretBox) EncryptSlot(dek []byte, label string) (*KeySlot, error) {
	var nonceSeed [24]byte
	if _, err := io.ReadFull(rand.Reader, nonceSeed[:]); err != nil {
		return nil, err
	}

	// return is basically append(nonceSeed, ciphertext...)
	ciphertext := secretbox.Seal([]byte{}, dek, deriveNonce(nonceSeed[:], label), &r.kek)

	dekEncrypted := append(nonceSeed[:], ciphertext...)

	return &KeySlot{
		Kind:         SlotKindNaclSecretBox,
		KekId:        r.KekId(),
		DekEncrypted: dekEncrypted,
	}, nil
}

func (r *naclSecretBox) DecryptSlot(slot *KeySlot, label string) ([]byte, error) {
	nonceSeed := slot.DekEncrypted[:24]

	// plaintext, ok := secretbox.Open(plaintext, e.EncryptedContent[24:], &nonce, &dekStatic)
	dek, ok := secretbox.Open([]byte{}, slot.DekEncrypted[24:], deriveNonce(nonceSeed[:], label), &r.kek)
	if !ok {
		return nil, errors.New("secretbox.Open failed")
	}

	return dek, nil
}

func (r *naclSecretBox) CanDecrypt(kind SlotKind, kekId string) bool {
	return kind == SlotKindNaclSecretBox && kekId == r.kekId
}

// By mixing nonce seed with label, we make label non-malleable
// nonce = first_24_bytes(sha256(nonceSeed || label))
func deriveNonce(nonceSeed []byte, label string) *[24]byte {
	// TODO: investigate why double append is necessary
	// hashInput := append(nonceSeed, []byte(label)...)
	hashInput := append(append([]byte{}, nonceSeed...), []byte(label)...)

	nonceAndLabelHashed := sha256.Sum256(hashInput)

	return to24(nonceAndLabelHashed[:24])
}

func to24(input []byte) *[24]byte {
	if len(input) != 24 {
		panic("to24 wrong len")
	}
	var output [24]byte
	copy(output[:], input[:24])
	return &output
}
