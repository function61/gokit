// Message Authentication Code - to make sure a message is not tampered with while
// in-transit by a malicious party. Small usability wrapper around crypto/hmac
package mac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

var ErrMacValidationFailed = errors.New("MAC validation failed")

type Mac struct {
	key     []byte
	message []byte
}

func New(key string, message string) *Mac {
	return &Mac{key: []byte(key), message: []byte(message)}
}

func (m *Mac) Authenticate(givenMac string) error {
	// safe from timing attacks
	if !hmac.Equal([]byte(givenMac), []byte(m.Sign())) {
		return ErrMacValidationFailed
	}

	return nil
}

// NOTE: entropy is lowered for prettier URLs
func (m *Mac) Sign() string {
	hmacHash := hmac.New(sha256.New, m.key)
	if _, err := hmacHash.Write(m.message); err != nil {
		panic(err)
	}

	sumB64 := base64.RawURLEncoding.EncodeToString(hmacHash.Sum(nil))

	// cap length for prettier URLs. we still have 96 bits of entropy
	return sumB64[0:16]
}
