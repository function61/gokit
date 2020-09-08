// Envelope encryption - envelope contains secret content encrypted with NaCl secretbox
// symmetric key ("DEK"), and that key is separately encrypted for each RSA public key recipient ("KEK").
package envelopeenc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	"github.com/function61/gokit/crypto/cryptoutil"
	"golang.org/x/crypto/nacl/secretbox"
)

// Protects its content with a DEK. the DEK is also contained in the envelope, but encrypted
// with multiple possible KEKs (public keys) as recipients, therefore the envelope can
// only be opened by any of the KEK private keys
type Envelope struct {
	KeySlots         []KeySlot `json:"key_slots"`
	Label            string    `json:"label,omitempty"` // non-malleable through RSA_OAEP_SHA256
	EncryptedContent []byte    `json:"content"`         // nonce || secretbox_ciphertext
}

// currently supports only one kind, but this is for extensibility
type SlotKind uint8

const (
	SlotKindRsaOaepSha256 SlotKind = 1
)

// The envelope is locked with several locks. Any of the KEKs can open the envelope
type KeySlot struct {
	Kind         SlotKind `json:"kind"`          // what kind of a key slot this is
	KekId        string   `json:"kek_id"`        // Kind=1 => SHA256-fingerprint of RSA public key
	DekEncrypted []byte   `json:"dek_encrypted"` // Kind=1 => RSA_OAEP_SHA256(kekPub, secretboxSecretKey, label)
}

type KekResolver func(kekId string) *rsa.PrivateKey

func Encrypt(plaintext []byte, keks []*rsa.PublicKey, label string) (*Envelope, error) {
	return encryptWithRand(plaintext, keks, label, rand.Reader)
}

func encryptWithRand(
	plaintext []byte,
	keks []*rsa.PublicKey,
	label string,
	cryptoRandReader io.Reader,
) (*Envelope, error) {
	var secretKey [32]byte
	if _, err := io.ReadFull(cryptoRandReader, secretKey[:]); err != nil {
		return nil, err
	}

	var nonce [24]byte
	if _, err := io.ReadFull(cryptoRandReader, nonce[:]); err != nil {
		return nil, err
	}

	keySlots := []KeySlot{}

	for _, kek := range keks {
		keySlot, err := makeKeySlot(secretKey[:], kek, label)
		if err != nil {
			return nil, err
		}

		keySlots = append(keySlots, *keySlot)
	}

	// return is basically append(nonce, ciphertext...)
	nonceAndCiphertext := secretbox.Seal(nonce[:], plaintext, &nonce, &secretKey)

	return &Envelope{
		KeySlots:         keySlots,
		Label:            label,
		EncryptedContent: nonceAndCiphertext,
	}, nil
}

func (e *Envelope) Decrypt(resolveKey KekResolver) ([]byte, error) {
	for _, slot := range e.KeySlots {
		// currently we only support one kind
		if slot.Kind != SlotKindRsaOaepSha256 {
			return nil, fmt.Errorf("unsupported slot kind: %d", slot.Kind)
		}

		if privKey := resolveKey(slot.KekId); privKey != nil {
			return e.decryptWithSlot(&slot, privKey)
		}
	}

	return nil, errors.New("no suitable private key found for any key slots")
}

func (e *Envelope) decryptWithSlot(slot *KeySlot, privKey *rsa.PrivateKey) ([]byte, error) {
	dek, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, slot.DekEncrypted, []byte(e.Label))
	if err != nil {
		return nil, fmt.Errorf("decryptWithSlot DecryptOAEP: %w", err)
	}

	var nonce [24]byte
	copy(nonce[:], e.EncryptedContent[:24])

	var dekStatic [32]byte
	copy(dekStatic[:], dek)

	plaintext := []byte{}
	plaintext, ok := secretbox.Open(plaintext, e.EncryptedContent[24:], &nonce, &dekStatic)
	if !ok {
		return nil, errors.New("secretbox.Open failed")
	}

	return plaintext, nil
}

func makeKeySlot(dek []byte, kekPub *rsa.PublicKey, label string) (*KeySlot, error) {
	kekId, err := cryptoutil.Sha256FingerprintForPublicKey(kekPub)
	if err != nil {
		return nil, err
	}

	dekCiphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, kekPub, dek, []byte(label))
	if err != nil {
		return nil, err
	}

	return &KeySlot{
		Kind:         SlotKindRsaOaepSha256,
		KekId:        kekId,
		DekEncrypted: dekCiphertext,
	}, nil
}

// helper for making a KEK resolver that only matches a single key that we possess
func SingleKey(privKey *rsa.PrivateKey) KekResolver {
	ourKekId, err := cryptoutil.Sha256FingerprintForPublicKey(&privKey.PublicKey)
	if err != nil {
		panic(err)
	}

	return func(slotKekId string) *rsa.PrivateKey {
		if slotKekId == ourKekId {
			return privKey
		} else {
			return nil
		}
	}
}
