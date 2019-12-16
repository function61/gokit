package cryptoutil

import (
	"crypto/x509"
	"io"
)

// PEM(cert)
func ParsePemX509Certificate(pemReader io.Reader) (*x509.Certificate, error) {
	certBytes, err := ParsePemBytes(pemReader, PemTypeCertificate)
	if err != nil {
		return nil, err
	}

	return x509.ParseCertificate(certBytes)
}

func Identity(cert x509.Certificate) string {
	if len(cert.DNSNames) > 0 {
		return cert.DNSNames[0]
	}

	return cert.Subject.CommonName
}

func Issuer(cert x509.Certificate) string {
	return cert.Issuer.Organization[0]
}
