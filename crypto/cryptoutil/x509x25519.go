package cryptoutil

// Go's stdlib x.509 PKCS #8 doesn't currently support X25519. OpenSSL does, so it's reason
// enough to replicate the data format against our tests.
// NOTE: This functionality is intended to be removed if/when stdlib gains X25519 support

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

// https://tools.ietf.org/html/rfc8410#section-3
var oidX25519 = asn1.ObjectIdentifier{1, 3, 101, 110}

func ParsePKCS8X25519PrivateKey(der []byte) (*X25519PrivateKey, error) {
	var privKey pkcs8
	if _, err := asn1.Unmarshal(der, &privKey); err != nil {
		return nil, fmt.Errorf("x509: %w", err)
	}

	if !privKey.Algo.Algorithm.Equal(oidX25519) {
		return nil, errors.New("x509: PKCS8 not oidX25519")
	}

	if privKey.Version != 0 {
		return nil, fmt.Errorf("x509: PKCS8 version not 0; got %d", privKey.Version)
	}

	var curvePrivateKey []byte
	if _, err := asn1.Unmarshal(privKey.PrivateKey, &curvePrivateKey); err != nil {
		return nil, fmt.Errorf("x509: invalid X25519 private key: %v", err)
	}

	if len(curvePrivateKey) != curve25519.ScalarSize {
		return nil, fmt.Errorf(
			"x509: invalid X25519 private key size: %d; expecting %d",
			len(privKey.PrivateKey),
			curve25519.ScalarSize)
	}

	return &X25519PrivateKey{curvePrivateKey}, nil
}

func ParsePemPKCS8X25519PrivateKey(input []byte) (*X25519PrivateKey, error) {
	der, err := ParsePemBytes(input, PemTypePrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parsePemPKCS8X25519PrivateKey: %w", err)
	}

	return ParsePKCS8X25519PrivateKey(der)
}

func MarshalPKCS8X25519PrivateKey(priv *X25519PrivateKey) ([]byte, error) {
	privateKey := priv.Bytes()
	if len(privateKey) != curve25519.ScalarSize {
		return nil, fmt.Errorf(
			"x509: invalid X25519 private key size: %d; expecting %d",
			len(privateKey),
			curve25519.ScalarSize)
	}

	privateKeyAsn1, err := asn1.Marshal(privateKey)
	if err != nil {
		return nil, fmt.Errorf("x509: marshalPKCS8X25519PrivateKey: %w", err)
	}

	return asn1.Marshal(pkcs8{
		Algo: pkix.AlgorithmIdentifier{
			Algorithm: oidX25519,
		},
		PrivateKey: privateKeyAsn1,
	})
}

func MarshalPemPKCS8X25519PrivateKey(priv *X25519PrivateKey) ([]byte, error) {
	der, err := MarshalPKCS8X25519PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	return MarshalPemBytes(der, PemTypePrivateKey), nil
}

func ParsePemPKCS8X25519PublicKey(pemBytes []byte) (*X25519PublicKey, error) {
	derBytes, err := ParsePemBytes(pemBytes, PemTypePublicKey)
	if err != nil {
		return nil, fmt.Errorf("x509: %w", err)
	}

	var pki publicKeyInfo
	if rest, err := asn1.Unmarshal(derBytes, &pki); err != nil {
		return nil, err
	} else if len(rest) != 0 {
		return nil, errors.New("x509: trailing data after ASN.1 of public-key")
	}

	if !pki.Algorithm.Algorithm.Equal(oidX25519) {
		return nil, errors.New("x509: unknown public key algorithm")
	}

	if pki.PublicKey.BitLength != curve25519.ScalarSize*8 {
		return nil, fmt.Errorf("x509: invalid BitLength for X25519 public key: %d", pki.PublicKey.BitLength)
	}

	return &X25519PublicKey{pki.PublicKey.Bytes}, nil
}

func marshalPKCS8X25519PublicKey(pubKey []byte) ([]byte, error) {
	if len(pubKey) != curve25519.ScalarSize {
		return nil, fmt.Errorf("x509: invalid length for X25519 public key: %d", len(pubKey))
	}

	return asn1.Marshal(publicKeyInfo{
		Algorithm: pkix.AlgorithmIdentifier{
			Algorithm: oidX25519,
		},
		PublicKey: asn1.BitString{
			Bytes:     pubKey,
			BitLength: len(pubKey) * 8,
		},
	})
}

func marshalPemPKCS8X25519PublicKey(pubKey X25519PublicKey) ([]byte, error) {
	der, err := marshalPKCS8X25519PublicKey(pubKey.Bytes())
	if err != nil {
		return nil, err
	}

	return MarshalPemBytes(der, PemTypePublicKey), nil
}

// pkcs8 reflects an ASN.1, PKCS #8 PrivateKey. See
// ftp://ftp.rsasecurity.com/pub/pkcs/pkcs-8/pkcs-8v1_2.asn
// and RFC 5208.
type pkcs8 struct {
	Version    int
	Algo       pkix.AlgorithmIdentifier
	PrivateKey []byte
	// optional attributes omitted.
}

type publicKeyInfo struct {
	Raw       asn1.RawContent
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}
