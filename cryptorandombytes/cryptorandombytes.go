// Safe random bytes in various formats (hex/base64/..)
package cryptorandombytes

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func Hex(bytesLen int) string {
	return hex.EncodeToString(randBytes(bytesLen))
}

func Base64Url(bytesLen int) string {
	return base64.RawURLEncoding.EncodeToString(randBytes(bytesLen))
}

// CLI arguments beginning with dash are problematic (which base64 URL variant can produce),
// so we'll be nice guys and guarantee that the ID won't start with one.
func Base64UrlWithoutLeadingDash(length int) string {
	id := Base64Url(length)

	if id[0] == '-' {
		// try again. the odds should exponentially decrease for recursion level to increase
		return Base64UrlWithoutLeadingDash(length)
	}

	return id
}

func randBytes(bytesLen int) []byte {
	bytes := make([]byte, bytesLen)

	// Read() uses io.ReadFull(), so non-nil error guarantees len(bytes) amount of bytes
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return bytes
}
