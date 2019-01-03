package cryptoutil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	PemTypeRsaPrivateKey = "RSA PRIVATE KEY"
	PemTypeRsaPublicKey  = "RSA PUBLIC KEY"
)

// PEM(PKCS1(rsa.PrivateKey))
func ParsePemPkcs1EncodedRsaPrivateKey(pemReader io.Reader) (*rsa.PrivateKey, error) {
	privKeyBytes, err := ParsePemBytes(pemReader, PemTypeRsaPrivateKey)
	if err != nil {
		return nil, err
	}

	privKey, err := x509.ParsePKCS1PrivateKey(privKeyBytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// PEM(PKCS1(rsa.PublicKey))
func ParsePemPkcs1EncodedRsaPublicKey(pemReader io.Reader) (*rsa.PublicKey, error) {
	pubKeyBytes, err := ParsePemBytes(pemReader, PemTypeRsaPublicKey)
	if err != nil {
		return nil, err
	}

	pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func MarshalPemBytes(content []byte, pemType string) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  pemType,
			Bytes: content,
		},
	)
}

func ParsePemBytes(pemReader io.Reader, expectedType string) ([]byte, error) {
	pemBytes, err := ioutil.ReadAll(pemReader)
	if err != nil {
		return nil, err
	}

	pemParsed, _ := pem.Decode(pemBytes)
	if pemParsed == nil {
		return nil, errors.New("PEM decode failed")
	}
	if pemParsed.Type != expectedType {
		return nil, fmt.Errorf(
			"unexpected PEM block type: %s; expecting %s",
			pemParsed.Type,
			expectedType)
	}

	return pemParsed.Bytes, nil
}
