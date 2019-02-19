package httpauth

import (
	"net/http"
	"time"
)

type RequestContext struct {
	User *UserDetails
}

type UserDetails struct {
	Id string
}

type HttpRequestAuthenticator interface {
	// authenticates a Request that has cookies returned by ToCookiesWithCsrfProtection()
	AuthenticateWithCsrfProtection(req *http.Request) (*UserDetails, error)
}

type Signer interface {
	Sign(userDetails UserDetails, now time.Time) string
}

// if returns nul, request handling is aborted.
// in return=nil case middleware is responsible for error response.
type MiddlewareChain func(w http.ResponseWriter, r *http.Request) *RequestContext

type MiddlewareChainMap map[string]MiddlewareChain
