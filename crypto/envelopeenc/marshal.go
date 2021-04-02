package envelopeenc

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

// func (e *EnvelopeBundle) Marshal(label string, recipients ...Recipient) ([]byte, []byte, error) {

func Marshal(label string, recipients ...Recipient) ([]byte, []byte, error) {
	bundle, dek, err := GenerateAndEncryptDek(label, recipients...)
	if err != nil {
		return nil, nil, err
	}

	bundleJson, err := json.Marshal(bundle)
	if err != nil {
		return nil, nil, err
	}

	headerMac, err := calculateHeaderMac(dek, bundleJson)
	if err != nil {
		return nil, nil, err
	}

	headerAndMac := fmt.Sprintf(
		"%s\n%s",
		bundleJson,
		headerMac)

	return []byte(headerAndMac), dek, nil
}

type Unmarshaled struct {
	// all fields private to force decryption via API that forces MAC validation

	bundle     EnvelopeBundle
	jsonToHash []byte
	claimedMac []byte
}

/*
func (e *EnvelopeBundle) DecryptWithResolver(resolveKek KekResolver) ([]byte, error) {
	for _, env := range e.Envelopes {
		if decrypter := resolveKek(env.Kind, env.KekId); decrypter != nil {
			return decrypter.DecryptSlot(&env, e.Label)
		}
	}

	return nil, errors.New("no suitable KEK found for any key slots")
}*/

func (u *Unmarshaled) Decrypt(decrypters ...Decrypter) ([]byte, error) {
	dek, err := u.bundle.Decrypt(decrypters...)
	if err != nil {
		return nil, err
	}

	// now that we know the key, we can calculate this
	expectedMac, err := calculateHeaderMac(dek, u.jsonToHash)
	if err != nil {
		return nil, fmt.Errorf("Decrypt: %w", err)
	}

	if subtle.ConstantTimeCompare(u.claimedMac, expectedMac) != 1 {
		return nil, errors.New("Decrypt: invalid MAC")
	}

	return dek, nil
}

func Unmarshal(marshaled []byte) (*Unmarshaled, error) {
	scanner := bufio.NewScanner(bytes.NewReader(marshaled))
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan first line: %w", scanner.Err())
	}
	jsonToHash := scanner.Bytes()

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan second line: %w", scanner.Err())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failure scanning second line: %w", scanner.Err())
	}

	claimedMac := scanner.Bytes()

	bundle := EnvelopeBundle{}
	if err := json.Unmarshal(jsonToHash, &bundle); err != nil {
		return nil, fmt.Errorf("failure parsing JSON: %w", err)
	}

	// TODO: assert for no extra data?

	return &Unmarshaled{
		bundle:     bundle,
		jsonToHash: jsonToHash,
		claimedMac: claimedMac,
	}, nil
}

func calculateHeaderMac(dek []byte, bundleJson []byte) ([]byte, error) {
	headerMacKey := make([]byte, 32)
	if _, err := io.ReadFull(hkdf.New(sha256.New, dek, nil, ourSalts.Header), headerMacKey); err != nil {
		return nil, err
	}

	headerMac := hmac.New(sha256.New, headerMacKey)
	if _, err := headerMac.Write(bundleJson); err != nil {
		return nil, err
	}

	return []byte(base64.RawURLEncoding.EncodeToString(headerMac.Sum(nil))), nil
}

// const marshalVersion = 1

/*	Format:

	uvarint  version (always 1)
	uvarint  length of EncryptedContent
	[]byte   EncryptedContent
	uvarint  amount of key slots

	for each key slot

		uvarint  length of KekId
		string   KekId
		uvarint  length of DekEncrypted
		[]byte   DekEncrypted

	NOTE: I would've gladly used Protobuf, but looks like you need a metric shit-ton of
	      imported (even runtime) code to use it.
*/

/*
// Marshals an envelope into a compact byte structure
// TODO: missing label and slot kind (DON'T USE)
func (e *Box) MarshalDONTUSE() ([]byte, error) {
	out := bytes.Buffer{}

	var err error

	writeBytes := func(data []byte) {
		if err != nil {
			return
		}

		_, err = out.Write(data)
	}

	writeUvarint := func(num uint64) {
		buf := make([]byte, binary.MaxVarintLen64)
		writeBytes(buf[0:binary.PutUvarint(buf, num)])
	}

	writeUvarint(uint64(marshalVersion))

	writeUvarint(uint64(len(e.EncryptedContent)))

	writeBytes(e.EncryptedContent)

	writeUvarint(uint64(len(e.KeySlots)))

	for _, keySlot := range e.KeySlots {
		writeUvarint(uint64(len(keySlot.KekId)))

		writeBytes([]byte(keySlot.KekId))

		writeUvarint(uint64(len(keySlot.DekEncrypted)))

		writeBytes(keySlot.DekEncrypted)
	}

	return out.Bytes(), err
}

func UnmarshalDONTUSE(buf []byte) (*Box, error) {
	bufReader := bytes.NewBuffer(buf)

	version, err := binary.ReadUvarint(bufReader)
	if err != nil {
		return nil, err
	}

	if version != marshalVersion {
		return nil, fmt.Errorf(
			"unexpected version: %d, expected %d",
			version,
			marshalVersion)
	}

	readByteSlice := func() ([]byte, error) {
		byteSliceLen, err := binary.ReadUvarint(bufReader)
		if err != nil {
			return nil, err
		}

		byteSlice := make([]byte, byteSliceLen)

		if _, err := io.ReadFull(bufReader, byteSlice); err != nil {
			return nil, err
		}

		return byteSlice, nil
	}

	encryptedContent, err := readByteSlice()
	if err != nil {
		return nil, err
	}

	keySlots := []KeySlot{}

	lenKeySlots, err := binary.ReadUvarint(bufReader)
	if err != nil {
		return nil, err
	}

	for i := uint64(0); i < lenKeySlots; i++ {
		kekId, err := readByteSlice()
		if err != nil {
			return nil, err
		}

		dekEncrypted, err := readByteSlice()
		if err != nil {
			return nil, err
		}

		keySlots = append(keySlots, KeySlot{
			KekId:        string(kekId),
			DekEncrypted: dekEncrypted,
		})
	}

	return &Box{
		EncryptedContent: encryptedContent,
		KeySlots:         keySlots,
	}, nil
}
*/
