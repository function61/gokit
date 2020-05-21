package httpauth

import (
	"net/http"

	"github.com/function61/gokit/csrf"
)

const (
	loginCookieName = "auth"
)

func ToCookiesWithCsrfProtection(tokenString string) []*http.Cookie {
	authCookie := &http.Cookie{
		Name:     loginCookieName,
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,                    // = not visible to JavaScript, to protect from XSS
		SameSite: http.SameSiteStrictMode, // CSRF protection
		Secure:   true,                    // only submit over https
	}

	return []*http.Cookie{
		authCookie,
		csrf.CreateCookie(),
	}
}

func ToCookie(tokenString string) *http.Cookie {
	return &http.Cookie{
		Name:     loginCookieName,
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,                    // = not visible to JavaScript, to protect from XSS
		SameSite: http.SameSiteStrictMode, // CSRF protection
		Secure:   true,                    // only submit over https
	}
}

func DeleteLoginCookie() *http.Cookie {
	// NOTE: keep cookie attributes in sync with ToCookie(), since the cookies may be
	//       considered separate cookies, unless components like "Path" (might be more) match
	return &http.Cookie{
		Name:     loginCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // => delete
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode, // CSRF protection
		Secure:   true,                    // only submit over https
	}
}
