// Use public key crypto to encrypt a data stream in a way, that the encrypting party can
// not necessarily decrypt the same data (unless she possesses the private key as well).
//
// DEPRECATED: this provides confidentiality, but is malleable (ciphertext is not authenticated)
//             use Age instead
package pkencryptedstream

/*	Public key (RSA) encrypted stream:

   - Producer can encrypt streams without knowing the consumer's private key, therefore
     producer loses access to the plaintext after the per-stream symmetric key leaves
     from memory.
   - Consumer can decrypt streams with the private key
   - Plaintext is encrypted with a symmetric per-stream random 256-bit AES key in CTR mode
   - Per-stream symmetric random AES key is public key -encrypted in an RSA-OAEP-SHA256 envelope
   - There is no authentication, so you don't know if ciphertext has been altered

        |  +-----------------------+
        |  |                       |
        |  | Below envelope length |
        |  |                       |
        |  +-----------------------+
        |
Stream  |  +-------------------------------+
        |  |                               |
        |  |  AES encryption key envelope  |
        |  |                               |
        |  | +---------------------------+ |
        |  | |                           | |
        |  | | RSA public key encryption | |
        |  | |                           | |
        |  | |   +-------------------+   | |
        |  | |   |                   |   | |
        |  | |   | AES symmetric key |   | |
        |  | |   |                   |   | |
        |  | |   +-------------------+   | |
        |  | |                           | |
        |  | +---------------------------+ |
        |  |                               |
        |  +-------------------------------+
        |
        |  +-----------------------+
        |  |                       |
        |  | IV, fixed # 16 bytes  |
        |  |                       |
        |  +-----------------------+
        |
        |  +--------------------------+
        |  |                          |
        |  | AES symmetric encryption |
        |  |                          |
        |  |      +-----------+       |
        |  |      |           |       |
        |  |      | Plaintext |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      |           |       |
        |  |      +-----------+       |
        v  |                          |
           +--------------------------+
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	aes256KeyBytes = 32
)

var (
	noLabel = []byte("")
	ivSize  = aes.BlockSize
)

// Format:
// 		<len of below envelope>
// 		<encrypted AES key envelope: RSA-OAEP-SHA256>
// 		<iv>
// 		<ciphertext stream>
func Writer(destination io.Writer, pubKey *rsa.PublicKey) (io.WriteCloser, error) {
	// generate new 256-bit AES key
	aesKey := make([]byte, aes256KeyBytes)
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}

	iv := make([]byte, ivSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	aesKeyEncrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, aesKey, noLabel)
	if err != nil {
		return nil, err
	}

	aesKeyEncryptedLen := len(aesKeyEncrypted)

	if err := binary.Write(destination, binary.LittleEndian, uint16(aesKeyEncryptedLen)); err != nil {
		return nil, err
	}

	if _, err := destination.Write(aesKeyEncrypted); err != nil {
		return nil, err
	}

	if _, err := destination.Write(iv); err != nil {
		return nil, err
	}

	aesBlockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	// more on mode of operation selection:
	// - https://crypto.stackexchange.com/questions/33846/is-regular-ctr-mode-vulnerable-to-any-attacks
	// - https://security.stackexchange.com/questions/27776/block-chaining-modes-to-avoid
	// - XTS would probably be the best, but that's not in the stdlib nor directly is an io.Writer
	return &cipher.StreamWriter{
		S: cipher.NewCTR(aesBlockCipher, iv),
		W: destination,
	}, nil
}

func Reader(source io.Reader, privateKey *rsa.PrivateKey) (io.Reader, error) {
	var aesKeyEncryptedLen uint16
	if err := binary.Read(source, binary.LittleEndian, &aesKeyEncryptedLen); err != nil {
		return nil, fmt.Errorf("reading aesKeyEncryptedLen: %v", err)
	}

	aesKeyEncrypted := make([]byte, aesKeyEncryptedLen)
	if _, err := io.ReadFull(source, aesKeyEncrypted); err != nil {
		return nil, fmt.Errorf("reading aesKeyEncrypted: %v", err)
	}

	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, aesKeyEncrypted, noLabel)
	if err != nil {
		return nil, fmt.Errorf("DecryptOAEP: %v", err)
	}

	aesBlockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, ivSize)
	if _, err := io.ReadFull(source, iv); err != nil {
		return nil, err
	}

	return &cipher.StreamReader{
		S: cipher.NewCTR(aesBlockCipher, iv),
		R: source,
	}, nil
}
