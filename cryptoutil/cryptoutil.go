// Small helpers for common crypto needs
package cryptoutil

import (
	"crypto"
	"crypto/ecdsa"
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
	PemTypeEcPrivateKey  = "EC PRIVATE KEY"
	PemTypeEcPublicKey   = "EC PUBLIC KEY"
	PemTypeCertificate   = "CERTIFICATE"
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

// PEM(cert)
func ParsePemX509Certificate(pemReader io.Reader) (*x509.Certificate, error) {
	certBytes, err := ParsePemBytes(pemReader, PemTypeCertificate)
	if err != nil {
		return nil, err
	}

	return x509.ParseCertificate(certBytes)
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

func ParsePemEncodedPrivateKey(serialized []byte) (crypto.PrivateKey, error) {
	block, _ := pem.Decode(serialized)
	if block == nil {
		return nil, errors.New("PEM decode failed")
	}

	switch block.Type {
	case PemTypeRsaPrivateKey:
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case PemTypeEcPrivateKey:
		return x509.ParseECPrivateKey(block.Bytes)
	default:
		return nil, fmt.Errorf("unknown private key type: %s", block.Type)
	}
}

func PublicKeyFromPrivateKey(priv crypto.PrivateKey) (crypto.PublicKey, error) {
	switch p := priv.(type) {
	case *rsa.PrivateKey:
		return &p.PublicKey, nil
	case *ecdsa.PrivateKey:
		return &p.PublicKey, nil
	default:
		return nil, errors.New("failed to get public key from private key")
	}
}

// "human readable" = don't ever try to parse the output format
func PublicKeyHumanReadableDescription(pubkey crypto.PublicKey) (string, error) {
	switch p := pubkey.(type) {
	case *rsa.PublicKey:
		return fmt.Sprintf("RSA-%d", p.Size()*8), nil
	case *ecdsa.PublicKey:
		return "ECDSA", nil
	default:
		return "", errors.New("unknown public key algorithm")
	}
}
