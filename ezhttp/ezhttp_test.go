package ezhttp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/function61/gokit/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

	_, err := Send(ctx, http.MethodGet, ts.URL)
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
	resp, err := Send(context.TODO(), http.MethodPost, ts.URL, SendJson(&req))
	assert.Assert(t, err == nil)
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), `{"Hello":"hello good sir"}`)
}

func TestPostArbitraryData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("received Content-Type: %s\n\n", r.Header.Get("Content-Type"))))

		if _, err := io.Copy(w, r.Body); err != nil { // pipe request to response
			panic(err)
		}
	}))
	defer ts.Close()

	reqBody := bytes.NewBufferString("why\nhello there\nmy good sir")
	resp, err := Send(context.TODO(), http.MethodPost, ts.URL, SendBody(reqBody, "text/awesome"))

	assert.Assert(t, err == nil)
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), `received Content-Type: text/awesome

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
	_, err := Send(context.TODO(), http.MethodGet, ts.URL+"/valid-json", RespondsJson(responseBody, false))
	assert.Assert(t, err == nil)
	assert.EqualString(t, responseBody.Hello, "World")

	_, err = Send(context.TODO(), http.MethodGet, ts.URL+"/valid-json-but-has-unknown-field", RespondsJson(responseBody, false))
	assert.EqualString(t, err.Error(), `json: unknown field "got"`)

	_, err = Send(context.TODO(), http.MethodGet, ts.URL+"/valid-json-but-has-unknown-field", RespondsJson(responseBody, true))
	assert.Assert(t, err == nil)
}

func TestRespondsJsonFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid-json")
	}))
	defer ts.Close()

	responseBody := &ExampleJsonPayload{}
	_, err := Send(context.TODO(), http.MethodGet, ts.URL, RespondsJson(responseBody, false))
	assert.EqualString(t, err.Error(), "invalid character 'i' looking for beginning of value")
}

func TestNon200x(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "I failed you :(", http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := Send(context.TODO(), http.MethodGet, ts.URL)
	assert.EqualString(t, err.Error(), "HTTP response not 2xx; was 500 Internal Server Error")

	_, err = Send(context.TODO(), http.MethodGet, ts.URL, TolerateNon2xxResponse)
	assert.Assert(t, err == nil)
}

func TestHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing User-Agent: %s", r.Header.Get("User-Agent"))
	}))
	defer ts.Close()

	resp, err := Send(context.TODO(), http.MethodGet, ts.URL, Header("User-Agent", "Sausage"))
	assert.Assert(t, err == nil)

	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), "Echoing User-Agent: Sausage")
}

func TestAuthBearer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing Authorization: %s", r.Header.Get("Authorization"))
	}))
	defer ts.Close()

	resp, err := Send(context.TODO(), http.MethodGet, ts.URL, AuthBearer("LOLOLOLOL"))
	assert.Assert(t, err == nil)

	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), "Echoing Authorization: Bearer LOLOLOLOL")
}

func TestAuthBasic(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echoing Authorization: %s", r.Header.Get("Authorization"))
	}))
	defer ts.Close()

	resp, err := Send(context.TODO(), http.MethodGet, ts.URL, AuthBasic("AzureDiamond", "hunter2"))
	assert.Assert(t, err == nil)

	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), "Echoing Authorization: Basic QXp1cmVEaWFtb25kOmh1bnRlcjI=")
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

	resp, err := Send(context.TODO(), http.MethodGet, ts.URL, Cookie(cmCookie))
	assert.Assert(t, err == nil)

	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.EqualString(t, string(respBody), `Echoing Cookie: cookiemonster="says nom nom"`)
}
