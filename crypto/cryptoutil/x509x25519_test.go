package cryptoutil

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestParsedPrivateKeyComputesToPublicKeyOpenSslEquivalent(t *testing.T) {
	bobPrivate, err := ParsePemPKCS8X25519PrivateKey([]byte(bobPrivatePemFromOpenssl))
	assert.Ok(t, err)

	bobPublicPem, err := marshalPemPKCS8X25519PublicKey(bobPrivate.Public())
	assert.Ok(t, err)

	// since we managed to import OpenSSL-exported X25519 key, computed its public key and
	// pubMarshaled equals what OpenSSL also computed from the private key, we know it works
	assert.Equal(t, string(bobPublicPem), bobPublicPemFromOpenssl)
}

func TestPrivateAndPublicKeyParsing(t *testing.T) {
	alicePrivate, err := ParsePemPKCS8X25519PrivateKey([]byte(alicePrivatePemFromOpenssl))
	assert.Ok(t, err)

	bobPrivate, err := ParsePemPKCS8X25519PrivateKey([]byte(bobPrivatePemFromOpenssl))
	assert.Ok(t, err)

	alicePublic := alicePrivate.Public()
	bobPublic, err := ParsePemPKCS8X25519PublicKey([]byte(bobPublicPemFromOpenssl))
	assert.Ok(t, err)

	assert.Assert(t, bytes.Equal(bobPublic.Bytes(), bobPrivate.Public().Bytes()))

	aliceShared, err := alicePrivate.ECDH(*bobPublic)
	assert.Ok(t, err)

	bobShared, err := bobPrivate.ECDH(alicePublic)
	assert.Ok(t, err)

	assert.Assert(t, bytes.Equal(aliceShared, bobShared))
	assert.Equal(t, fmt.Sprintf("%x", aliceShared), "1e9be8d5c01de22df9ccd852c17db8ca9367f72140cc1bbed9bac9e0ee7c1e53")

	/*
		Ideally we would like to compare OpenSSL's ECDH result to in-Go result (thus verifying
		our private/public were parsed correctly), but OpenSSL's "pkeyutl -derive" doesn't seem
		to be raw ECDH shared secret.

		// generated using:
		//   $ openssl pkeyutl -derive -out alicebob.key -inkey alice.key -peerkey bob.pub
		openSslDeriveResult := []byte{
			0x9b, 0x1e, 0xd5, 0xe8, 0x1d, 0xc0, 0x2d, 0xe2, 0xcc, 0xf9, 0x52, 0xd8, 0x7d, 0xc1, 0xca, 0xb8,
			0x67, 0x93, 0x21, 0xf7, 0xcc, 0x40, 0xbe, 0x1b, 0xba, 0xd9, 0xe0, 0xc9, 0x7c, 0xee, 0x53, 0x1e,
		}
	*/
}

const (
	// $ openssl genpkey -algorithm X25519 -out alice.key
	alicePrivatePemFromOpenssl = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VuBCIEIJgKLwy+fvq4bWzu0DnyXO6suA7ydA8ix0/oYYv9ttJQ
-----END PRIVATE KEY-----`
	// $ openssl genpkey -algorithm X25519 -out bob.key
	bobPrivatePemFromOpenssl = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VuBCIEIIAgiLS1r1ZkXWIC0iVtiWHXaEaM+9BB/xpVlYMnuf1t
-----END PRIVATE KEY-----`
	// $ openssl pkey -pubout -in bob.key -out bob.pub
	bobPublicPemFromOpenssl = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VuAyEA4xvsubO9QWVwAeWGv9/QBRwT6MWY8GvO7Y/MF8oH+ns=
-----END PUBLIC KEY-----
`
)
