// If you know the hash of a stream beforehand, wrap it with this to get
// integrity-verified io.Reader
package hashverifyreader

import (
	"bytes"
	"errors"
	"hash"
	"io"
)

/*
    +--------+
    |        |
    | Source |
    |        |
    +---+----+
        |
        |               +------+
      +-v---+           |      |
      |     +-----------> Hash |
      | Tee |           |      |
      |     |           +---+--+
      +--+--+               |
         |           +------v-----+
         |           |            |
         |    +------+ HashVerify |
1. until |    |      |            |
 tee EOF |    |      +------------+
         |    | 2. after tee
      +--v----v-----+ EOF
      |             |
      | MultiReader |
      |             |
      +------+------+
             |
             |
        +----v-----+
        |          |
        | Consumer |
        |          |
        +----------+
*/

// creates a two-part reader whose first part is connected to source, and the second part
// (= when source read to EOF) Read() (on hashVerifyReader) immediately returns EOF if hash matches
func New(source io.Reader, hash hash.Hash, expectedHash []byte) io.Reader {
	// create a T-junction to pump the read bytes 1:1 to hash fn
	sourceTee := io.TeeReader(source, hash)

	// when we get EOF from the Tee, last write is guaranteed to be pumped to the hash.
	// MultiReader is read sequentially so when Read() is finally called on our
	// hashVerifyReader we can safely access the finished digest and can return error to
	// the consumer if the source data was corrupted
	return io.MultiReader(sourceTee, &hashVerifyReader{
		expectedHash: expectedHash,
		hash:         hash,
	})
}

type hashVerifyReader struct {
	expectedHash []byte
	hash         hash.Hash
}

func (h *hashVerifyReader) Read([]byte) (int, error) {
	if !bytes.Equal(h.hash.Sum(nil), h.expectedHash) {
		return 0, errors.New("hashVerifyReader: digest mismatch")
	}

	return 0, io.EOF
}
