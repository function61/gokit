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

func randBytes(bytesLen int) []byte {
	bytes := make([]byte, bytesLen)

	// Read() uses io.ReadFull(), so non-nil error guarantees len(bytes) amount of bytes
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return bytes
}
