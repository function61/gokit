package cryptoutil

import (
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestMarshalPemBytes(t *testing.T) {
	assert.Equal(t, string(MarshalPemBytes([]byte("hello"), PemTypeCertificate)), `-----BEGIN CERTIFICATE-----
aGVsbG8=
-----END CERTIFICATE-----
`)
}

func TestParsePemPkcs1EncodedRsaPrivateKey(t *testing.T) {
	privKey, err := ParsePemPkcs1EncodedRsaPrivateKey([]byte(testValidRsaPrivateKey))
	assert.Ok(t, err)
	assert.Equal(t, privKey.E, 65537)
}

func TestParsePemPkcs1EncodedRsaPublicKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPublicKey([]byte(testValidRsaPublicKey))
	assert.Ok(t, err)
	assert.Equal(t, pubKey.E, 65537)

	// invalid PEM
	_, err = ParsePemPkcs1EncodedRsaPublicKey([]byte("-----BEGIN RSA PUBLIC KEY-----\nMIIB@\n-----END RSA PUBLIC KEY-----"))
	assert.Equal(t, err.Error(), "PEM decode failed")

	// valid PEM, invalid asn1
	_, err = ParsePemPkcs1EncodedRsaPublicKey([]byte("-----BEGIN RSA PUBLIC KEY-----\nMIIB\n-----END RSA PUBLIC KEY-----"))
	assert.Equal(t, err.Error(), "asn1: syntax error: truncated tag or length")
}

func TestParsePemEncodedPrivateKey(t *testing.T) {
	rsaKey, err := ParsePemEncodedPrivateKey([]byte(testValidRsaPrivateKey))
	assert.Ok(t, err)
	rsaPublicKey, _ := PublicKeyFromPrivateKey(rsaKey)
	rsaPublicKeyDescr, _ := PublicKeyHumanReadableDescription(rsaPublicKey)
	assert.Equal(t, rsaPublicKeyDescr, "RSA-1024")

	ecdsaKey, err := ParsePemEncodedPrivateKey([]byte(testValidEcdsaPrivateKey))
	assert.Ok(t, err)
	ecdsaPublicKey, _ := PublicKeyFromPrivateKey(ecdsaKey)
	ecdsaPublicKeyDescr, _ := PublicKeyHumanReadableDescription(ecdsaPublicKey)
	assert.Equal(t, ecdsaPublicKeyDescr, "ECDSA")
}

func TestSha256FingerprintForPublicKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPublicKey([]byte(testValidRsaPublicKey))
	assert.Ok(t, err)

	fingerprint, err := Sha256FingerprintForPublicKey(pubKey)
	assert.Ok(t, err)

	assert.Equal(t, fingerprint, "SHA256:mkKvVUiFg0oZd1IRltdabZPDD4oPUJcyIFnxeqi1sl8")
}

func TestMarshalPemPkcs1EncodedRsaPublicKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPublicKey([]byte(testValidRsaPublicKey))
	assert.Ok(t, err)

	assert.Equal(t, string(MarshalPemPkcs1EncodedRsaPublicKey(pubKey)), `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA+xGZ/wcz9ugFpP07Nspo6U17l0YhFiFpxxU4pTk3Lifz9R3zsIsu
ERwta7+fWIfxOo208ett/jhskiVodSEt3QBGh4XBipyWopKwZ93HHaDVZAALi/2A
+xTBtWdEo7XGUujKDvC2/aZKukfjpOiUI8AhLAfjmlcD/UZ1QPh0mHsglRNCmpCw
mwSXA9VNmhz+PiB+Dml4WWnKW/VHo2ujTXxq7+efMU4H2fny3Se3KYOsFPFGZ1TN
QSYlFuShWrHPtiLmUdPoP6CV2mML1tk+l7DIIqXrQhLUKDACeM5roMx0kLhUWB8P
+0uj1CNlNN4JRZlC7xFfqiMbFRU9Z4N6YwIDAQAB
-----END RSA PUBLIC KEY-----
`)
}

func TestMarshalPemPkcs1EncodedRsaPrivateKey(t *testing.T) {
	pubKey, err := ParsePemPkcs1EncodedRsaPrivateKey([]byte(testValidRsaPrivateKey))
	assert.Ok(t, err)

	assert.Equal(t, string(MarshalPemPkcs1EncodedRsaPrivateKey(pubKey)), `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZf
XJ/DgYSF6vUpwmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/
3j+skZ6UtW+5u09lHNsj6tQ51s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQAB
AoGAFijko56+qGyN8M0RVyaRAXz++xTqHBLh3tx4VgMtrQ+WEgCjhoTwo23KMBAu
JGSYnRmoBZM3lMfTKevIkAidPExvYCdm5dYq3XToLkkLv5L2pIIVOFMDG+KESnAF
V7l2c+cnzRMW0+b6f8mR1CJzZuxVLL6Q02fvLi55/mbSYxECQQDeAw6fiIQXGukB
I4eMZZt4nscy2o12KyYner3VpoeE+Np2q+Z3pvAMd/aNzQ/W9WaI+NRfcxUJrmfP
wIGm63ilAkEAxCL5HQb2bQr4ByorcMWm/hEP2MZzROV73yF41hPsRC9m66KrheO9
HPTJuo3/9s5p+sqGxOlFL0NDt4SkosjgGwJAFklyR1uZ/wPJjj611cdBcztlPdqo
xssQGnh85BzCj/u3WqBpE2vjvyyvyI5kX6zk7S0ljKtt2jny2+00VsBerQJBAJGC
1Mg5Oydo5NwD6BiROrPxGo2bpTbu/fhrT8ebHkTz2eplU9VQQSQzY1oZMVX8i1m5
WUTLPz2yLJIBQVdXqhMCQBGoiuSoSjafUhV7i1cEGpb88h5NBYZzWXGZ37sJ5QsW
+sJyoNde3xH8vdXhzU7eT82D6X/scw9RZz+/6rCJ4p0=
-----END RSA PRIVATE KEY-----
`)
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
MIICXAIBAAKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZf
XJ/DgYSF6vUpwmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/
3j+skZ6UtW+5u09lHNsj6tQ51s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQAB
AoGAFijko56+qGyN8M0RVyaRAXz++xTqHBLh3tx4VgMtrQ+WEgCjhoTwo23KMBAu
JGSYnRmoBZM3lMfTKevIkAidPExvYCdm5dYq3XToLkkLv5L2pIIVOFMDG+KESnAF
V7l2c+cnzRMW0+b6f8mR1CJzZuxVLL6Q02fvLi55/mbSYxECQQDeAw6fiIQXGukB
I4eMZZt4nscy2o12KyYner3VpoeE+Np2q+Z3pvAMd/aNzQ/W9WaI+NRfcxUJrmfP
wIGm63ilAkEAxCL5HQb2bQr4ByorcMWm/hEP2MZzROV73yF41hPsRC9m66KrheO9
HPTJuo3/9s5p+sqGxOlFL0NDt4SkosjgGwJAFklyR1uZ/wPJjj611cdBcztlPdqo
xssQGnh85BzCj/u3WqBpE2vjvyyvyI5kX6zk7S0ljKtt2jny2+00VsBerQJBAJGC
1Mg5Oydo5NwD6BiROrPxGo2bpTbu/fhrT8ebHkTz2eplU9VQQSQzY1oZMVX8i1m5
WUTLPz2yLJIBQVdXqhMCQBGoiuSoSjafUhV7i1cEGpb88h5NBYZzWXGZ37sJ5QsW
+sJyoNde3xH8vdXhzU7eT82D6X/scw9RZz+/6rCJ4p0=
-----END RSA PRIVATE KEY-----`
	testValidEcdsaPrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEINTwTX5Xt26rkBv44y2dXEwetcT54HZr6v20FBFhW7hboAoGCCqGSM49
AwEHoUQDQgAExER1ikX4DnifPRJ1i9Ra9H9IGq7CaKbBFgofAo5c63l+pYd+/7Zr
K8BX/2/x4tBJR1w8YyKepbbwQtX3zVI7mw==
-----END EC PRIVATE KEY-----`
)
