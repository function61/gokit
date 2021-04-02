// Envelope encryption - envelope contains secret content encrypted with NaCl secretbox
// symmetric key ("DEK"), and that key is separately encrypted for each RSA public key recipient ("KEK").
package envelopeenc

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

type Recipient interface {
	KekId() string
	EncryptSlot(dek []byte, label string) (*Envelope, error)
}

type Decrypter interface {
	KekId() string
	DecryptSlot(slot *Envelope, label string) ([]byte, error)
	// CanDecrypt(kind RecipientKind, kekId string) bool
}

// Protects its content with a DEK. the DEK is also contained in the envelope, but encrypted
// with multiple possible KEKs (public keys) as recipients, therefore the envelope can
// only be opened by any of the KEK private keys
type EnvelopeBundle struct {
	Recipients map[string]EnvelopeRaw `json:"recip"`           // keyed by recipient's KEK ID
	Label      string                 `json:"label,omitempty"` // non-malleable through constructions in each RecipientKind
}

type EnvelopeRaw []byte

/*
type RecipientKind uint8

const (
	RecipientKindX25519        RecipientKind = 1 // X25519, anchored to Age
	RecipientKindElliptic      RecipientKind = 2 // P-256/P-384/P-512
	RecipientKindSymmetric     RecipientKind = 3
	RecipientKindPassword      RecipientKind = 4 // scrypt, anchored to Age
	RecipientKindRsaOaepSha256 RecipientKind = 5 // RSA, anchored to Age
)
*/

// The envelope is locked with several locks. Any of the KEKs can open the envelope
type Envelope struct {
	// Kind         RecipientKind `json:"kind"`
	KekId        string `json:"recip"`
	DekEncrypted []byte `json:"env"`
}

// type KekResolver func(kind RecipientKind, kekId string) Decrypter

// generates & encrypts a DEK
func GenerateAndEncryptDek(
	label string,
	recipients ...Recipient,
) (*EnvelopeBundle, []byte, error) {
	dek := make([]byte, chacha20poly1305.KeySize)
	if _, err := io.ReadFull(rand.Reader, dek); err != nil {
		return nil, nil, err
	}

	recipientEnvelopes := map[string]EnvelopeRaw{}

	for _, kek := range recipients {
		env, err := kek.EncryptSlot(dek, label)
		if err != nil {
			return nil, nil, err
		}

		recipientEnvelopes[kek.KekId()] = env.DekEncrypted
	}

	// prevent accidents from resulting in effectively lost data
	if len(recipientEnvelopes) == 0 {
		return nil, nil, errors.New("envelope with zero recipients not supported")
	}

	return &EnvelopeBundle{
		Recipients: recipientEnvelopes,
		Label:      label,
	}, dek, nil
}

func (e *EnvelopeBundle) Decrypt(decrypters ...Decrypter) ([]byte, error) {
	for _, decrypter := range decrypters {
		if recipient, found := e.Recipients[decrypter.KekId()]; found {
			return decrypter.DecryptSlot(&Envelope{
				DekEncrypted: recipient,
			}, e.Label)
		}
	}

	return nil, errors.New("no suitable KEK found for any key slots")

	// return e.DecryptWithResolver(NewKekResolver(decrypters...))
}

/*
func (e *EnvelopeBundle) DecryptWithResolver(resolveKek KekResolver) ([]byte, error) {
	for kekId, env := range e.Recipients {
		if decrypter := resolveKek(env.Kind, kekId); decrypter != nil {
			return decrypter.DecryptSlot(&env, e.Label)
		}
	}

	return nil, errors.New("no suitable KEK found for any key slots")
}

func NewKekResolver(decrypters ...Decrypter) KekResolver {
	return func(kind RecipientKind, kekId string) Decrypter {
		for _, decrypter := range decrypters {
			if decrypter.CanDecrypt(kind, kekId) {
				return decrypter
			}
		}

		return nil
	}
}
*/

func concatSlices(slices ...[]byte) []byte {
	concat := []byte{}

	for _, slice := range slices {
		concat = append(concat, slice...)
	}

	return concat
}
