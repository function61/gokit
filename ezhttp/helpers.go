package ezhttp

import (
	"crypto/tls"
	"net/http"
)

const (
	jsonContentType = "application/json"
)

var InsecureTlsClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

// checks if "err" is *ResponseStatusError and has "statusCode" status
func ErrorIs(err error, statusCode int) bool {
	if err, isStatusError := err.(*ResponseStatusError); isStatusError && err.StatusCode() == statusCode {
		return true
	} else {
		return false
	}
}
