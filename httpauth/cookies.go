package httpauth

import (
	"net/http"

	"github.com/function61/gokit/csrf"
)

const (
	loginCookieName = "login"
)

func ToCookiesWithCsrfProtection(tokenString string) []*http.Cookie {
	authCookie := &http.Cookie{
		Name:     loginCookieName,
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true, // = not visible to JavaScript, to protect from XSS
		// Secure: true, // FIXME
	}

	return []*http.Cookie{
		authCookie,
		csrf.CreateCookie(),
	}
}

func DeleteLoginCookie() *http.Cookie {
	// NOTE: keep cookie attributes in sync with ToCookie(), since the cookies may be
	//       considered separate cookies, unless components like "Path" (might be more) match
	return &http.Cookie{
		Name:     loginCookieName,
		Value:    "del",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // => delete
		// Secure: true, // FIXME
	}
}
