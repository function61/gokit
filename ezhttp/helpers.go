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
