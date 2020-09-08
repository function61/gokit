package hashverifyreader

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestReaderGoodHash(t *testing.T) {
	input := bytes.NewBufferString("The quick brown fox jumps over the lazy dog")

	verifiedReader := New(
		input,
		sha256.New(),
		hexToBytes("d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"))

	out, err := ioutil.ReadAll(verifiedReader)
	assert.Ok(t, err)
	assert.EqualString(t, string(out), "The quick brown fox jumps over the lazy dog")
}

func TestReaderBadHash(t *testing.T) {
	input := bytes.NewBufferString("The quick brown fox jumps over the lazy dog")

	thisHashWillNotMatch := New(
		input,
		sha256.New(),
		hexToBytes("7885198ae80b48ebf6c0d0d6b8ce86169570d346ec0b2c43029d04b3e3763ec3"))

	_, err := ioutil.ReadAll(thisHashWillNotMatch)
	assert.EqualString(t, err.Error(), "hashVerifyReader: digest mismatch")
}

func hexToBytes(input string) []byte {
	out, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	return out
}
