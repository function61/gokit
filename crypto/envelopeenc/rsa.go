package envelopeenc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/function61/gokit/crypto/cryptoutil"
)

type rsaPublicKey struct {
	pubKey *rsa.PublicKey
	kekId  string
}

var _ SlotEncrypter = (*rsaPublicKey)(nil)

func RsaOaepSha256Encrypter(pubKey *rsa.PublicKey) *rsaPublicKey {
	kekId, err := cryptoutil.Sha256FingerprintForPublicKey(pubKey)
	if err != nil {
		panic(err)
	}

	return newRsaOaepSha256Encrypter(pubKey, kekId)
}

func newRsaOaepSha256Encrypter(pubKey *rsa.PublicKey, kekId string) *rsaPublicKey {
	return &rsaPublicKey{
		pubKey: pubKey,
		kekId:  kekId,
	}
}

func (r *rsaPublicKey) KekId() string {
	return r.kekId
}

func (r *rsaPublicKey) EncryptSlot(dek []byte, label string) (*KeySlot, error) {
	dekCiphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, r.pubKey, dek, []byte(label))
	if err != nil {
		return nil, err
	}

	return &KeySlot{
		Kind:         SlotKindRsaOaepSha256,
		KekId:        r.KekId(),
		DekEncrypted: dekCiphertext,
	}, nil
}

type rsaPrivateKey struct {
	*rsaPublicKey // private key also knows its public key, so it can go in both directions
	privKey       *rsa.PrivateKey
}

var _ SlotDecrypter = (*rsaPrivateKey)(nil)
var _ SlotEncrypter = (*rsaPrivateKey)(nil)

func RsaOaepSha256Decrypter(privKey *rsa.PrivateKey) *rsaPrivateKey {
	kekId, err := cryptoutil.Sha256FingerprintForPublicKey(&privKey.PublicKey)
	if err != nil {
		panic(err)
	}

	return &rsaPrivateKey{
		rsaPublicKey: newRsaOaepSha256Encrypter(&privKey.PublicKey, kekId),
		privKey:      privKey,
	}
}

func (r *rsaPrivateKey) DecryptSlot(slot *KeySlot, label string) ([]byte, error) {
	dek, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, r.privKey, slot.DekEncrypted, []byte(label))
	if err != nil {
		return nil, fmt.Errorf("decryptWithSlot DecryptOAEP: %w", err)
	}

	return dek, nil
}

func (r *rsaPrivateKey) CanDecrypt(kind SlotKind, kekId string) bool {
	return kind == SlotKindRsaOaepSha256 && kekId == r.kekId
}
