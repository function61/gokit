// Envelope encryption - envelope contains secret content encrypted with NaCl secretbox
// symmetric key ("DEK"), and that key is separately encrypted for each RSA public key recipient ("KEK").
package envelopeenc

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

type SlotEncrypter interface {
	EncryptSlot(dek []byte, label string) (*KeySlot, error)
}

type SlotDecrypter interface {
	DecryptSlot(slot *KeySlot, label string) ([]byte, error)
	CanDecrypt(kind SlotKind, kekId string) bool
}

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
	SlotKindNaclSecretBox SlotKind = 2
)

// The envelope is locked with several locks. Any of the KEKs can open the envelope
type KeySlot struct {
	Kind         SlotKind `json:"kind"`          // what kind of a key slot this is
	KekId        string   `json:"kek_id"`        // Kind=1 => SHA256-fingerprint of RSA public key, Kind=2 => name for key
	DekEncrypted []byte   `json:"dek_encrypted"` // Kind=1 => RSA_OAEP_SHA256(kekPub, dek, label), Kind=2 => nonceSeed || naclSecretBoxSeal(dek, deriveNonce(nonceSeed, label), kek)
}

type KekResolver func(kind SlotKind, kekId string) SlotDecrypter

func Encrypt(plaintext []byte, slotEncrypters []SlotEncrypter, label string) (*Envelope, error) {
	return encryptWithRand(plaintext, slotEncrypters, label, rand.Reader)
}

func encryptWithRand(
	plaintext []byte,
	slotEncrypters []SlotEncrypter,
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

	for _, kek := range slotEncrypters {
		keySlot, err := kek.EncryptSlot(secretKey[:], label)
		if err != nil {
			return nil, err
		}

		keySlots = append(keySlots, *keySlot)
	}

	if len(keySlots) == 0 {
		return nil, errors.New("envelope with zero slots not supported")
	}

	// return is basically append(nonce, ciphertext...)
	nonceAndCiphertext := secretbox.Seal(nonce[:], plaintext, &nonce, &secretKey)

	return &Envelope{
		KeySlots:         keySlots,
		Label:            label,
		EncryptedContent: nonceAndCiphertext,
	}, nil
}

func (e *Envelope) Decrypt(decrypters ...SlotDecrypter) ([]byte, error) {
	return e.DecryptWithResolver(NewKekResolver(decrypters...))
}

// use this when you want to manage resolving the KEK yourself
func (e *Envelope) DecryptWithResolver(resolveKek KekResolver) ([]byte, error) {
	for _, slot := range e.KeySlots {
		if decrypter := resolveKek(slot.Kind, slot.KekId); decrypter != nil {
			slot := slot // pin

			dek, err := decrypter.DecryptSlot(&slot, e.Label)
			if err != nil {
				return nil, err
			}

			return e.decryptData(dek)
		}
	}

	return nil, errors.New("no suitable KEK found for any key slots")
}

func (e *Envelope) decryptData(dek []byte) ([]byte, error) {
	var nonce [24]byte
	copy(nonce[:], e.EncryptedContent[:24])

	var dekStatic [32]byte
	copy(dekStatic[:], dek)

	plaintext, ok := secretbox.Open(nil, e.EncryptedContent[24:], &nonce, &dekStatic)
	if !ok {
		return nil, errors.New("secretbox.Open failed")
	}

	return plaintext, nil
}

func NewKekResolver(decrypters ...SlotDecrypter) KekResolver {
	return func(kind SlotKind, kekId string) SlotDecrypter {
		for _, decrypter := range decrypters {
			if decrypter.CanDecrypt(kind, kekId) {
				return decrypter
			}
		}

		return nil
	}
}
