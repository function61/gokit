package cryptoutil

// Safe random bytes in various formats (hex/base64/..)

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func RandHex(bytesLen int) string {
	return hex.EncodeToString(cryptoRandBytes(bytesLen))
}

func RandBase64Url(bytesLen int) string {
	return base64.RawURLEncoding.EncodeToString(cryptoRandBytes(bytesLen))
}

// CLI arguments beginning with dash are problematic (which base64 URL variant can produce),
// so this variant is for nice guys to make life simpler (reduces entropy very little)
func RandBase64UrlWithoutLeadingDash(bytesLen int) string {
	id := RandBase64Url(bytesLen)

	if id[0] == '-' {
		// try again. the odds should exponentially decrease for recursion level to increase
		return RandBase64UrlWithoutLeadingDash(bytesLen)
	}

	return id
}

func cryptoRandBytes(bytesLen int) []byte {
	bytes := make([]byte, bytesLen)

	// Read() uses io.ReadFull(), so non-nil error guarantees len(bytes) amount of bytes
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return bytes
}
