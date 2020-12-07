package cryptoutil

// Deprecated - will be removed soon.

import (
	"crypto/rsa"
	"crypto/x509"
)

// PEM(PKCS1(rsa.PrivateKey))
func ParsePemPkcs1EncodedRsaPrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	privKeyBytes, err := ParsePemBytes(pemBytes, PemTypeRsaPrivateKey)
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
func ParsePemPkcs1EncodedRsaPublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	pubKeyBytes, err := ParsePemBytes(pemBytes, PemTypeRsaPublicKey)
	if err != nil {
		return nil, err
	}

	pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func MarshalPemPkcs1EncodedRsaPublicKey(pubKey *rsa.PublicKey) []byte {
	return MarshalPemBytes(x509.MarshalPKCS1PublicKey(pubKey), PemTypeRsaPublicKey)
}

func MarshalPemPkcs1EncodedRsaPrivateKey(privKey *rsa.PrivateKey) []byte {
	return MarshalPemBytes(x509.MarshalPKCS1PrivateKey(privKey), PemTypeRsaPrivateKey)
}
