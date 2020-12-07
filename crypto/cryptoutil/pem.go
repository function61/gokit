package cryptoutil

import (
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	PemTypePrivateKey    = "PRIVATE KEY" // PKCS #8 (= supports different types via parametrization)
	PemTypePublicKey     = "PUBLIC KEY"  // PKCS #8
	PemTypeRsaPrivateKey = "RSA PRIVATE KEY"
	PemTypeRsaPublicKey  = "RSA PUBLIC KEY"
	PemTypeEcPrivateKey  = "EC PRIVATE KEY"
	PemTypeEcPublicKey   = "EC PUBLIC KEY"
	PemTypeCertificate   = "CERTIFICATE"
)

func MarshalPemBytes(content []byte, pemType string) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  pemType,
			Bytes: content,
		},
	)
}

func ParsePemBytes(pemBytes []byte, expectedType string) ([]byte, error) {
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
