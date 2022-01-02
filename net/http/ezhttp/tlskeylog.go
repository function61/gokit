package ezhttp

// To facilitate HTTPS packet inspection in Wireshark etc.
// Great concrete instructions: https://wiki.wireshark.org/TLS
//
// (Go doesn't have automatic support for SSLKEYLOGFILE, instead we need explicit opt-in.)

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var (
	keyLoggingClientCached   *http.Client
	keyLoggingClientCachedMu sync.Mutex
)

// opt-in for logging TLS secrets to file specified in SSLKEYLOGFILE ENV var.
// if logging is not requested, uses http.DefaultClient
var EnableTLSKeyLog = After(func(conf *Config) {
	keyLoggingClientCachedMu.Lock()
	defer keyLoggingClientCachedMu.Unlock()

	if keyLoggingClientCached == nil { // build only once
		// should this be on by default or not:
		// > for most people it represents an easy way for a local actor to compromise their cryptographic keys by silently logging them
		// https://bugzilla.mozilla.org/show_bug.cgi?id=1188657
		// > I think that SSLKEYLOGFILE support should be enabled for Debian builds of libnss, since otherwise users who need this functionality have to make their own builds of libnss.
		// https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=842292
		keyLogWriter, err := func() (io.Writer, error) {
			if name := os.Getenv("SSLKEYLOGFILE"); name != "" {
				return os.Create(name)
			} else {
				return nil, nil
			}
		}()
		if err != nil {
			// other options would be for our function to return error or to define a roundtripper that always fails
			panic(fmt.Errorf("SSLKEYLOGFILE requested but: %w", err))
		}

		keyLoggingClientCached = func() *http.Client {
			if keyLogWriter != nil {
				return &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							KeyLogWriter: keyLogWriter,
						},
					},
				}
			} else {
				return http.DefaultClient
			}
		}()
	}

	// now use either:
	// a) http.DefaultClient or
	// b) our instance with keyLogWriter specified
	conf.Client = keyLoggingClientCached
})
