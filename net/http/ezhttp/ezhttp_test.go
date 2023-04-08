package ezhttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
)

// shared across many tests
type ExampleJsonPayload struct {
	Hello string
}

func TestTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Millisecond)
	defer cancel()

	_, err := Get(ctx, ts.URL)
	assert.Matches(t, err.Error(), "Get [^ ]+ context deadline exceeded")
}

func TestPostJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			panic("no JSON Content-Type specified")
		}

		if _, err := io.Copy(w, r.Body); err != nil { // pipe request to response
			panic(err)
		}
	}))
	defer ts.Close()

	req := ExampleJsonPayload{Hello: "hello good sir"}
	resp, err := Post(context.TODO(), ts.URL, SendJson(&req))
	assert.Ok(t, err)
	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), `{"Hello":"hello good sir"}`)
}

func TestPostArbitraryData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(fmt.Sprintf("received Content-Type: %s\n\n", r.Header.Get("Content-Type")))); err != nil {
			panic(err)
		}

		if _, err := io.Copy(w, r.Body); err != nil { // pipe request to response
			panic(err)
		}
	}))
	defer ts.Close()

	reqBody := bytes.NewBufferString("why\nhello there\nmy good sir")
	resp, err := Post(context.TODO(), ts.URL, SendBody(reqBody, "text/awesome"))

	assert.Ok(t, err)
	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), `received Content-Type: text/awesome

why
hello there
my good sir`)
}

func TestRespondsJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/valid-json":
			fmt.Fprintln(w, `{"Hello": "World"}`)
		case "/valid-json-but-has-unknown-field":
			fmt.Fprintln(w, `{"Hello": "World", "got": "more than you asked for"}`)
		default:
			panic("invalid path")
		}
	}))
	defer ts.Close()

	responseBody := &ExampleJsonPayload{}
	_, err := Get(context.TODO(), ts.URL+"/valid-json", RespondsJson(responseBody, false))
	assert.Ok(t, err)
	assert.Equal(t, responseBody.Hello, "World")

	_, err = Get(context.TODO(), ts.URL+"/valid-json-but-has-unknown-field", RespondsJson(responseBody, false))
	assert.Equal(t, err.Error(), `json: unknown field "got"`)

	_, err = Get(context.TODO(), ts.URL+"/valid-json-but-has-unknown-field", RespondsJson(responseBody, true))
	assert.Ok(t, err)
}

func TestRespondsJsonFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid-json")
	}))
	defer ts.Close()

	responseBody := &ExampleJsonPayload{}
	_, err := Get(context.TODO(), ts.URL, RespondsJson(responseBody, false))
	assert.Equal(t, err.Error(), "invalid character 'i' looking for beginning of value")
}

func TestNon200x(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "I failed you :(", http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := Get(context.TODO(), ts.URL)
	assert.Equal(t, err.Error(), "500 Internal Server Error; I failed you :(\n")

	_, err = Get(context.TODO(), ts.URL, TolerateNon2xxResponse)
	assert.Ok(t, err)
}

func TestHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing User-Agent: %s", r.Header.Get("User-Agent"))
	}))
	defer ts.Close()

	resp, err := Get(context.TODO(), ts.URL, Header("User-Agent", "Sausage"))
	assert.Ok(t, err)

	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), "Echoing User-Agent: Sausage")
}

func TestAuthBearer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing Authorization: %s", r.Header.Get("Authorization"))
	}))
	defer ts.Close()

	resp, err := Get(context.TODO(), ts.URL, AuthBearer("LOLOLOLOL"))
	assert.Ok(t, err)

	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), "Echoing Authorization: Bearer LOLOLOLOL")
}

func TestRespondsJSONDoesntOverrideExplicitAccept(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"AcceptEchoed": "%s"}`, r.Header.Get("Accept"))
	}))
	defer ts.Close()

	respJSON := struct {
		AcceptEchoed string
	}{}

	// before the fix, RespondsJSONDisallowUnknownFields() used to override explicit header
	_, err := Get(context.TODO(), ts.URL, Header("Accept", "text/foobar"), RespondsJSONDisallowUnknownFields(&respJSON))
	assert.Ok(t, err)

	assert.Equal(t, respJSON.AcceptEchoed, "text/foobar")
}

func TestAuthBasic(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing Authorization: %s", r.Header.Get("Authorization"))
	}))
	defer ts.Close()

	resp, err := Get(context.TODO(), ts.URL, AuthBasic("AzureDiamond", "hunter2"))
	assert.Ok(t, err)

	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), "Echoing Authorization: Basic QXp1cmVEaWFtb25kOmh1bnRlcjI=")
}

func TestCookie(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing Cookie: %s", r.Header.Get("Cookie"))
	}))
	defer ts.Close()

	cmCookie := http.Cookie{
		Name:  "cookiemonster",
		Value: "says nom nom",
	}

	resp, err := Get(context.TODO(), ts.URL, Cookie(cmCookie))
	assert.Ok(t, err)

	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, string(respBody), `Echoing Cookie: cookiemonster="says nom nom"`)
}

func TestExpectedVsUnexpectedNotModified(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotModified)
	}))
	defer ts.Close()

	_, err := Get(context.TODO(), ts.URL)
	assert.Equal(t, err.Error(), "304 Not Modified; <no response body>")

	// it's not an error however, if "expecting caching" headers are sent

	_, err = Get(context.TODO(), ts.URL, Header("If-None-Match", `"myCoolETag"`))
	assert.Ok(t, err)
}

func TestErrorIs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "I failed you :(", http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := Get(context.TODO(), ts.URL)

	assert.Equal(t, ErrorIs(err, http.StatusInternalServerError), true)
	assert.Equal(t, ErrorIs(err, http.StatusNotModified), false)
	assert.Equal(t, ErrorIs(nil, http.StatusNotModified), false)
}

func TestRequestBodyForNonBodyMethods(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	_, err := Get(context.TODO(), ts.URL, SendBody(strings.NewReader("huh?"), "text/plain"))
	assert.Equal(t, err.Error(), "ezhttp: GET with non-nil body is usually a mistake")

	_, err = Head(context.TODO(), ts.URL, SendBody(strings.NewReader("huh?"), "text/plain"))
	assert.Equal(t, err.Error(), "ezhttp: HEAD with non-nil body is usually a mistake")
}
