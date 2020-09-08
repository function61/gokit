package httpauth

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/function61/gokit/testing/assert"
)

func TestSignAndAuthenticate(t *testing.T) {
	signer, _ := NewEcJwtSigner([]byte(testPrivateKey))

	token := signer.Sign(UserDetails{Id: "123"}, time.Now())

	bothCookies := ToCookiesWithCsrfProtection(token)
	assert.Assert(t, len(bothCookies) == 2)

	assert.EqualString(t, bothCookies[0].Name, "auth")
	assert.EqualString(t, bothCookies[0].Value, token)
	assert.Assert(t, bothCookies[0].HttpOnly)
	assert.EqualString(t, bothCookies[1].Name, "csrf_token")
	assert.Assert(t, len(bothCookies[1].Value) == 22)
	assert.Assert(t, !bothCookies[1].HttpOnly)

	authenticator, _ := NewEcJwtAuthenticator([]byte(testPublicKey))

	authenticateReq := func(req *http.Request) string {
		userDetails, err := authenticator.AuthenticateWithCsrfProtection(req)

		if err == nil {
			// cannot print whole JWT token because it contains random data for crypto
			return fmt.Sprintf("userid<%s> tok<%s>", userDetails.Id, userDetails.AuthTokenJwt[0:8]+"..")
		}

		return err.Error()
	}

	onlyLoginCookie := []*http.Cookie{bothCookies[0]}
	onlyCsrfCookie := []*http.Cookie{bothCookies[1]}

	assert.EqualString(t,
		authenticateReq(makeReq(bothCookies, "invalid header")),
		"csrf: cookie does not match HTTP header")

	assert.EqualString(t,
		authenticateReq(makeReq(bothCookies, "")),
		"csrf: x-csrf-token HTTP header missing")

	assert.EqualString(t,
		authenticateReq(makeReq(onlyLoginCookie, "just something here")),
		"csrf: cookie csrf_token missing")

	assert.EqualString(t,
		authenticateReq(makeReq(onlyCsrfCookie, "just something here")),
		"csrf: cookie does not match HTTP header")

	assert.EqualString(t,
		authenticateReq(makeReq(bothCookies, onlyCsrfCookie[0].Value)),
		"userid<123> tok<eyJhbGci..>")

	assert.EqualString(t,
		authenticateReq(makeReq(onlyCsrfCookie, onlyCsrfCookie[0].Value)),
		"auth: either specify 'auth' cookie or 'Authorization' header")

	// authenticate via header instead of cookie

	reqWithBearerToken := makeReq(onlyCsrfCookie, onlyCsrfCookie[0].Value)
	reqWithBearerToken.Header.Set("Authorization", "Bearer "+onlyLoginCookie[0].Value)

	assert.EqualString(t, authenticateReq(reqWithBearerToken), "userid<123> tok<eyJhbGci..>")
}

func TestSignAndAuthenticateMismatchingPublicKey(t *testing.T) {
	signer, _ := NewEcJwtSigner([]byte(testPrivateKey))
	// this public key is not linked to the private key
	authenticator, _ := NewEcJwtAuthenticator([]byte(testMismatchingPublicKey))

	token := signer.Sign(UserDetails{Id: "123"}, time.Now())

	bothCookies := ToCookiesWithCsrfProtection(token)

	_, err := authenticator.AuthenticateWithCsrfProtection(makeReq(bothCookies, bothCookies[1].Value))

	assert.EqualString(t, err.Error(), "JWT authentication: crypto/ecdsa: verification error")
}

func TestTokenExpiry(t *testing.T) {
	signer, _ := NewEcJwtSigner([]byte(testPrivateKey))
	authenticator, _ := NewEcJwtAuthenticator([]byte(testPublicKey))

	t0 := time.Date(2019, 2, 19, 15, 0, 0, 0, time.UTC)

	token := signer.Sign(UserDetails{Id: "123"}, t0)
	bothCookies := ToCookiesWithCsrfProtection(token)

	shouldBeValid := func(should bool) {
		t.Helper()
		userDetails, err := authenticator.AuthenticateWithCsrfProtection(makeReq(bothCookies, bothCookies[1].Value))

		if should {
			assert.Ok(t, err)
			assert.EqualString(t, userDetails.Id, "123")
		} else {
			assert.EqualString(t, err.Error(), "JWT authentication: token is expired by 1h0m0s")
		}
	}

	timewarp(t0, func() {
		shouldBeValid(true)
	})

	timewarp(t0.Add(8*time.Hour), func() {
		shouldBeValid(true)
	})

	timewarp(t0.Add(12*time.Hour), func() {
		shouldBeValid(true)
	})

	timewarp(t0.Add(23*time.Hour), func() {
		shouldBeValid(true)
	})

	timewarp(t0.Add(25*time.Hour), func() {
		shouldBeValid(false)
	})
}

// you cannot pass time to jwt-go, but instead we must change the global function. brilliant,
// as different time could never be used in production due to race conditions..
func timewarp(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func makeReq(cookies []*http.Cookie, csrfToken string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://dummy/", nil)
	if csrfToken != "" {
		req.Header.Set("x-csrf-token", csrfToken)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	return req
}

const (
	testPrivateKey           = "-----BEGIN PRIVATE KEY-----\nMIHcAgEBBEIB2evM5cezr4xQNCmYdcCma1u3jwkprSvYp58Pil8+Zzb2CNjKiZrF\n7RJZpd/lbBazd4B35avQJxhZAye0Ri9j9kCgBwYFK4EEACOhgYkDgYYABAG9B1oJ\nu7Vt0LdDSm0acZgywfsdSCRbxsSxju821nxBAee2kg4CcXVKd2166scow3wS3+Ve\n2O+JgduwXtS2DRLwzgHrrxqv6CkHhXbP2FXAxAzN+R1/9+121pQ3yB2pSWX2nxvS\nawExMIUPbfdnR20V6CmArvdf5WjQtnaBVtC9kLTLXg==\n-----END PRIVATE KEY-----\n"
	testPublicKey            = "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBvQdaCbu1bdC3Q0ptGnGYMsH7HUgk\nW8bEsY7vNtZ8QQHntpIOAnF1SndteurHKMN8Et/lXtjviYHbsF7Utg0S8M4B668a\nr+gpB4V2z9hVwMQMzfkdf/ftdtaUN8gdqUll9p8b0msBMTCFD233Z0dtFegpgK73\nX+Vo0LZ2gVbQvZC0y14=\n-----END PUBLIC KEY-----\n"
	testMismatchingPublicKey = "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBRevOMH3BS0QMj3/BbNmEOLs2xTgF\nRiC2CEDYX5Gw1sgPCMwRaBdtxXycad5MlpoHnK69Mr4am0MDrijnNrO3oM8BeqaO\nElvMsPvg8G9auizjg17WeqikULV+n9nDk1od0AuYIvfkwh0e3hp1u6WkYA1ppyMZ\ny+PtZBkj7Q327dT07BY=\n-----END PUBLIC KEY-----\n"
)
