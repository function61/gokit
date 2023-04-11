// This package aims to wrap Go HTTP Client's request-response with sane defaults:
//
//   - You are forced to consider timeouts by having to specify Context
//   - Instead of not considering non-2xx status codes as a failure, check that by default
//     (unless explicitly asked to)
//   - Sending and receiving JSON requires much less boilerplate, and on receiving JSON you
//     are forced to think whether to "allowUnknownFields"
package ezhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	DefaultTimeout10s = 10 * time.Second
	NoOpConfig        = ConfigPiece{} // sometimes it's beneficial to give an option that does nothing (so user doesn't have to do varargs)
)

type ConfigHook func(conf *Config)

func After(fn ConfigHook) ConfigPiece {
	return ConfigPiece{AfterInit: fn}
}

// same as After(), but Config.Request is nil. used mainly for specifying
// request body, which must be known on call to http.NewRequest()
func Before(fn ConfigHook) ConfigPiece {
	return ConfigPiece{BeforeInit: fn}
}

type ConfigPiece struct {
	BeforeInit ConfigHook
	AfterInit  ConfigHook
}

type Config struct {
	Abort                         error // ConfigHook can set this to abort request send
	Client                        *http.Client
	Request                       *http.Request
	TolerateNon2xxResponse        bool
	RequestBody                   io.Reader
	OutputsJson                   bool
	OutputsJsonRef                interface{}
	OutputsJsonAllowUnknownFields bool
}

type ResponseStatusError struct {
	error
	statusCode int
}

// returns the (non-2xx) status code that caused the error
func (e ResponseStatusError) StatusCode() int {
	return e.statusCode
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func Get(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return newRequest(ctx, http.MethodGet, url, confPieces...).Send()
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func Post(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return newRequest(ctx, http.MethodPost, url, confPieces...).Send()
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func Put(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return newRequest(ctx, http.MethodPut, url, confPieces...).Send()
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func Head(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return newRequest(ctx, http.MethodHead, url, confPieces...).Send()
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func Del(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return newRequest(ctx, http.MethodDelete, url, confPieces...).Send()
}

func newRequest(ctx context.Context, method string, url string, confPieces ...ConfigPiece) *Config {
	conf := &Config{
		Client: http.DefaultClient,
	}

	withErr := func(err error) *Config {
		conf.Abort = err // will be early-error-returned in `Send()`
		return conf
	}

	for _, configure := range confPieces {
		if configure.BeforeInit == nil {
			continue
		}
		configure.BeforeInit(conf)
	}

	if conf.Abort != nil {
		return withErr(conf.Abort)
	}

	// "Request has body = No" for:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD
	if conf.RequestBody != nil && (method == http.MethodGet || method == http.MethodHead) {
		// Technically, these can have body, but it's usually a mistake so if we need it we'll
		// make it an opt-in flag.
		return withErr(fmt.Errorf("ezhttp: %s with non-nil body is usually a mistake", method))
	}

	req, err := http.NewRequest(
		method,
		url,
		conf.RequestBody)
	if err != nil {
		return withErr(err)
	}

	req = req.WithContext(ctx)

	conf.Request = req

	for _, configure := range confPieces {
		if configure.AfterInit == nil {
			continue
		}
		configure.AfterInit(conf)
	}

	return conf
}

func (conf *Config) Send() (*http.Response, error) {
	if conf.Abort != nil {
		return nil, conf.Abort
	}

	resp, err := conf.Client.Do(conf.Request)
	if err != nil {
		return resp, err // this is a transport-level error
	}

	// 304 is an error unless caller is expecting such response by sending caching headers
	if resp.StatusCode == http.StatusNotModified && conf.Request.Header.Get("If-None-Match") != "" {
		return resp, nil
	}

	// handle application-level errors
	if !conf.TolerateNon2xxResponse && (resp.StatusCode < 200 || resp.StatusCode > 299) {
		defer resp.Body.Close()

		// TODO: if caller wants to process error herself, we need an opt-out for this mechanism
		return resp, errorWithResponseBodySample(resp)
	}

	if conf.OutputsJson {
		defer resp.Body.Close()

		jsonDecoder := json.NewDecoder(resp.Body)
		if !conf.OutputsJsonAllowUnknownFields {
			jsonDecoder.DisallowUnknownFields()
		}

		if err := jsonDecoder.Decode(conf.OutputsJsonRef); err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func errorWithResponseBodySample(resp *http.Response) error {
	errContentSampleLength := 128
	truncatedIndicator := ""

	// .Body is documented as always non-nil
	errContent, err := io.ReadAll(io.LimitReader(resp.Body, int64(errContentSampleLength)))
	if err != nil {
		errContent = []byte(fmt.Sprintf("<failed reading response body: %v>", err))
	} else if len(errContent) == errContentSampleLength {
		truncatedIndicator = ".."
	}

	if len(errContent) == 0 {
		errContent = []byte("<no response body>")
	}

	return &ResponseStatusError{
		statusCode: resp.StatusCode,
		error:      fmt.Errorf("%s; %s%s", resp.Status, errContent, truncatedIndicator),
	}
}
