package ezhttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
)

const (
	jsonContentType = "application/json"
)

var InsecureTlsClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			//nolint:gosec // ok as this is explicit opt-in
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

// same as the corresponding without "New" prefix, but just prepared the request configuration without sending it yet
func NewGet(ctx context.Context, url string, confPieces ...ConfigPiece) *Config {
	return newRequest(ctx, http.MethodGet, url, confPieces...)
}

// same as the corresponding without "New" prefix, but just prepared the request configuration without sending it yet
func NewPost(ctx context.Context, url string, confPieces ...ConfigPiece) *Config {
	return newRequest(ctx, http.MethodPost, url, confPieces...)
}

// same as the corresponding without "New" prefix, but just prepared the request configuration without sending it yet
func NewPut(ctx context.Context, url string, confPieces ...ConfigPiece) *Config {
	return newRequest(ctx, http.MethodPut, url, confPieces...)
}

// same as the corresponding without "New" prefix, but just prepared the request configuration without sending it yet
func NewHead(ctx context.Context, url string, confPieces ...ConfigPiece) *Config {
	return newRequest(ctx, http.MethodHead, url, confPieces...)
}

// same as the corresponding without "New" prefix, but just prepared the request configuration without sending it yet
func NewDel(ctx context.Context, url string, confPieces ...ConfigPiece) *Config {
	return newRequest(ctx, http.MethodDelete, url, confPieces...)
}

// for `method` please use `net/http` "enum" (quotes because it's not declared as such)
func (c *Config) CURLEquivalent() ([]string, error) {
	if err := c.Abort; err != nil {
		return nil, err
	}

	req := c.Request // shorthand

	cmd := []string{"curl", "--request=" + req.Method}

	for key, values := range req.Header {
		// FIXME: doesn't take into account multiple values
		cmd = append(cmd, fmt.Sprintf("--header=%s=%s", key, values[0]))
	}

	cmd = append(cmd, req.URL.String())

	return cmd, nil
}
