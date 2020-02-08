package cryptoutil

import (
	"github.com/function61/gokit/assert"
	"strings"
	"testing"
)

func TestMarshalPemBytes(t *testing.T) {
	assert.EqualString(t, string(MarshalPemBytes([]byte("hello"), PemTypeCertificate)), `-----BEGIN CERTIFICATE-----
aGVsbG8=
-----END CERTIFICATE-----
`)
}

func TestParsePemPkcs1EncodedRsaPrivateKey(t *testing.T) {
	privKey, err := ParsePemPkcs1EncodedRsaPrivateKey(strings.NewReader(testValidRsaPrivateKey))
	assert.Ok(t, err)
	assert.Assert(t, privKey.E == 65537)
}

func TestParsePemPkcs1EncodedRsaPublicKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPublicKey(strings.NewReader(testValidRsaPublicKey))
	assert.Ok(t, err)
	assert.Assert(t, pubKey.E == 65537)

	// invalid PEM
	_, err = ParsePemPkcs1EncodedRsaPublicKey(strings.NewReader("-----BEGIN RSA PUBLIC KEY-----\nMIIB@\n-----END RSA PUBLIC KEY-----"))
	assert.EqualString(t, err.Error(), "PEM decode failed")

	// valid PEM, invalid asn1
	_, err = ParsePemPkcs1EncodedRsaPublicKey(strings.NewReader("-----BEGIN RSA PUBLIC KEY-----\nMIIB\n-----END RSA PUBLIC KEY-----"))
	assert.EqualString(t, err.Error(), "asn1: syntax error: truncated tag or length")
}

func TestParsePemEncodedPrivateKey(t *testing.T) {
	rsaKey, err := ParsePemEncodedPrivateKey([]byte(testValidRsaPrivateKey))
	assert.Ok(t, err)
	rsaPublicKey, _ := PublicKeyFromPrivateKey(rsaKey)
	rsaPublicKeyDescr, _ := PublicKeyHumanReadableDescription(rsaPublicKey)
	assert.EqualString(t, rsaPublicKeyDescr, "RSA-1024")

	ecdsaKey, err := ParsePemEncodedPrivateKey([]byte(testValidEcdsaPrivateKey))
	assert.Ok(t, err)
	ecdsaPublicKey, _ := PublicKeyFromPrivateKey(ecdsaKey)
	ecdsaPublicKeyDescr, _ := PublicKeyHumanReadableDescription(ecdsaPublicKey)
	assert.EqualString(t, ecdsaPublicKeyDescr, "ECDSA")
}

func TestSha256FingerprintForPublicKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPublicKey(strings.NewReader(testValidRsaPublicKey))
	assert.Ok(t, err)

	fingerprint, err := Sha256FingerprintForPublicKey(pubKey)
	assert.Ok(t, err)

	assert.EqualString(t, fingerprint, "SHA256:mkKvVUiFg0oZd1IRltdabZPDD4oPUJcyIFnxeqi1sl8")
}

const (
	testValidRsaPublicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA+xGZ/wcz9ugFpP07Nspo6U17l0YhFiFpxxU4pTk3Lifz9R3zsIsu
ERwta7+fWIfxOo208ett/jhskiVodSEt3QBGh4XBipyWopKwZ93HHaDVZAALi/2A
+xTBtWdEo7XGUujKDvC2/aZKukfjpOiUI8AhLAfjmlcD/UZ1QPh0mHsglRNCmpCw
mwSXA9VNmhz+PiB+Dml4WWnKW/VHo2ujTXxq7+efMU4H2fny3Se3KYOsFPFGZ1TN
QSYlFuShWrHPtiLmUdPoP6CV2mML1tk+l7DIIqXrQhLUKDACeM5roMx0kLhUWB8P
+0uj1CNlNN4JRZlC7xFfqiMbFRU9Z4N6YwIDAQAB
-----END RSA PUBLIC KEY-----`
	testValidRsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZfXJ/DgYSF6vUp
wmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/3j+skZ6UtW+5u09lHNsj6tQ5
1s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQABAoGAFijko56+qGyN8M0RVyaRAXz++xTqHBLh
3tx4VgMtrQ+WEgCjhoTwo23KMBAuJGSYnRmoBZM3lMfTKevIkAidPExvYCdm5dYq3XToLkkLv5L2
pIIVOFMDG+KESnAFV7l2c+cnzRMW0+b6f8mR1CJzZuxVLL6Q02fvLi55/mbSYxECQQDeAw6fiIQX
GukBI4eMZZt4nscy2o12KyYner3VpoeE+Np2q+Z3pvAMd/aNzQ/W9WaI+NRfcxUJrmfPwIGm63il
AkEAxCL5HQb2bQr4ByorcMWm/hEP2MZzROV73yF41hPsRC9m66KrheO9HPTJuo3/9s5p+sqGxOlF
L0NDt4SkosjgGwJAFklyR1uZ/wPJjj611cdBcztlPdqoxssQGnh85BzCj/u3WqBpE2vjvyyvyI5k
X6zk7S0ljKtt2jny2+00VsBerQJBAJGC1Mg5Oydo5NwD6BiROrPxGo2bpTbu/fhrT8ebHkTz2epl
U9VQQSQzY1oZMVX8i1m5WUTLPz2yLJIBQVdXqhMCQBGoiuSoSjafUhV7i1cEGpb88h5NBYZzWXGZ
37sJ5QsW+sJyoNde3xH8vdXhzU7eT82D6X/scw9RZz+/6rCJ4p0=
-----END RSA PRIVATE KEY-----`
	testValidEcdsaPrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINTwTX5Xt26rkBv44y2dXEwetcT54HZr6v20FBFhW7hboAoGCCqGSM49
AwEHoUQDQgAExER1ikX4DnifPRJ1i9Ra9H9IGq7CaKbBFgofAo5c63l+pYd+/7Zr
K8BX/2/x4tBJR1w8YyKepbbwQtX3zVI7mw==
-----END EC PRIVATE KEY-----`
)
