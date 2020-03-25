// CSRF protection
package csrf

// Using a double-submit cookie here. We cannot just double-submit the session cookie,
// because we mark it as HttpOnly to prevent an XSS attacker from stealing it.
//
// We use a stateless, random string (not deriving the CSRF token from the secret). The user
// can in theory change it from her browser, but that's ok because while an attacker could
// inject the CSRF header, the attacker cannot set a cookie. The whole security model of this
// relies on the attacker not being able to set a cookie for another domain in the forged request:
//    https://stackoverflow.com/questions/6761415/how-to-set-a-cookie-for-another-domain
//    https://developer.mozilla.org/en-US/docs/Glossary/Forbidden_header_name
//
// This is thoroughly tested by httpauth package

import (
	"errors"
	"net/http"

	"github.com/function61/gokit/cryptorandombytes"
)

const (
	csrfCookieName = "csrf_token"
	csrfHeaderName = "x-csrf-token"
)

var (
	errHeaderMissing           = errors.New("csrf: " + csrfHeaderName + " HTTP header missing")
	errCookieMissing           = errors.New("csrf: cookie " + csrfCookieName + " missing")
	errCookieAndHeaderMismatch = errors.New("csrf: cookie does not match HTTP header")
)

func CreateCookie() *http.Cookie {
	return &http.Cookie{
		Name:     csrfCookieName,
		Value:    cryptorandombytes.Base64Url(16),
		Path:     "/",
		HttpOnly: false, // = visible to JavaScript (has to be)
		// Secure: true, // TODO
	}
}

func Validate(req *http.Request) error {
	headerToken := req.Header.Get(csrfHeaderName)
	if headerToken == "" {
		return errHeaderMissing
	}

	cookie, err := req.Cookie(csrfCookieName)
	if err != nil {
		return errCookieMissing
	}

	if headerToken != cookie.Value {
		return errCookieAndHeaderMismatch
	}

	return nil
}
