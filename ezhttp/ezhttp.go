// This package aims to wrap Go HTTP Client's request-response with sane defaults:
//
// - You are forced to consider timeouts by having to specify Context
// - Instead of not considering non-2xx status codes as a failure, check that by default
//   (unless explicitly asked to)
// - Sending and receiving JSON requires much less boilerplate, and on receiving JSON you
//   are forced to think whether to "allowUnknownFields"
package ezhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	DefaultTimeout10s = 10 * time.Second
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

func Get(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return do(ctx, http.MethodGet, url, confPieces...)
}

func Post(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return do(ctx, http.MethodPost, url, confPieces...)
}

func Put(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return do(ctx, http.MethodPut, url, confPieces...)
}

func Head(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return do(ctx, http.MethodHead, url, confPieces...)
}

func Del(ctx context.Context, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	return do(ctx, http.MethodDelete, url, confPieces...)
}

// returns *ResponseStatusError as error if non-2xx response (unless TolerateNon2xxResponse()).
// error is not *ResponseStatusError for transport-level errors, content (JSON) marshaling errors etc
func do(ctx context.Context, method string, url string, confPieces ...ConfigPiece) (*http.Response, error) {
	conf := &Config{
		Client: http.DefaultClient,
	}

	for _, configure := range confPieces {
		if configure.BeforeInit == nil {
			continue
		}
		configure.BeforeInit(conf)
	}

	if conf.Abort != nil {
		return nil, conf.Abort
	}

	req, err := http.NewRequest(
		method,
		url,
		conf.RequestBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	conf.Request = req

	for _, configure := range confPieces {
		if configure.AfterInit == nil {
			continue
		}
		configure.AfterInit(conf)
	}

	if conf.Abort != nil {
		return nil, conf.Abort
	}

	resp, err := conf.Client.Do(req)
	if err != nil {
		return resp, err // this is a transport-level error
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
	errContent, err := ioutil.ReadAll(io.LimitReader(resp.Body, int64(errContentSampleLength)))
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
