package cryptoutil

import (
	"fmt"

	"golang.org/x/crypto/curve25519"
)

// Wrapper around Curve25519's X25519 to not just pass around []byte blobs and thus to prevent
// mistakes

type X25519PublicKey struct {
	key []byte
}

func NewX25519PublicKey(key []byte) X25519PublicKey {
	if len(key) != curve25519.PointSize {
		panic(fmt.Errorf("X25519PrivateKey: incorrect public key size: %d", len(key)))
	}

	return X25519PublicKey{key}
}

func (p X25519PublicKey) Bytes() []byte {
	return p.key
}

type X25519PrivateKey struct {
	key []byte
}

func NewX25519PrivateKey(b []byte) (*X25519PrivateKey, error) {
	if len(b) != curve25519.ScalarSize {
		return nil, fmt.Errorf("X25519PrivateKey: incorrect private key size: %d", len(b))
	}

	return &X25519PrivateKey{b}, nil
}

func (p X25519PrivateKey) Bytes() []byte {
	return p.key
}

func (p X25519PrivateKey) ECDH(peer X25519PublicKey) ([]byte, error) {
	return curve25519.X25519(p.key, peer.key)
}

func (p X25519PrivateKey) Public() X25519PublicKey {
	publicKey, err := curve25519.X25519(p.key, curve25519.Basepoint)
	if err != nil {
		panic(err)
	}

	return X25519PublicKey{key: publicKey}
}
