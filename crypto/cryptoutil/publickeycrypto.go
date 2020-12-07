// Small helpers for common crypto needs
package cryptoutil

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
)

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
	case PemTypePrivateKey: // PKCS #8
		return x509.ParsePKCS8PrivateKey(block.Bytes)
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
	case *ed25519.PrivateKey:
		return p.Public(), nil
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
	case *ed25519.PublicKey:
		return "Ed25519", nil
	default:
		return "", errors.New("unknown public key algorithm")
	}
}

func Sha256FingerprintForPublicKey(publicKey crypto.PublicKey) (string, error) {
	// need to convert to ssh.PublicKey to be able to use the fingerprint util
	sshPubKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	return ssh.FingerprintSHA256(sshPubKey), nil
}
