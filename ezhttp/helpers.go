package ezhttp

import (
	"crypto/tls"
	"net/http"
)

var InsecureTlsClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}
