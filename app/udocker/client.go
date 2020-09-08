// μDocker - super-lightweight Docker client. pulling in the official client uses
// huge amounts of code and memory
package udocker

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/function61/gokit/os/osutil"
)

type certificateLoader func() (*tls.Certificate, error)

func Client(
	dockerUrl string,
	clientCertificateLoader certificateLoader,
	insecureSkipVerify bool,
) (*http.Client, string, error) {
	u, err := url.Parse(dockerUrl)
	if err != nil {
		return nil, "", err
	}

	if u.Scheme == "unix" { // unix socket needs own dialer
		return &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
					dialer := net.Dialer{} // don't know why we need a struct to use DialContext()
					return dialer.DialContext(ctx, "unix", u.Path)
				},
			},
		}, "http://localhost", nil
	}

	clientCertificate, err := clientCertificateLoader()
	if err != nil {
		return nil, "", err
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*clientCertificate},
		InsecureSkipVerify: insecureSkipVerify,
	}

	return &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}, dockerUrl, nil
}

func ClientCertificateFromEnv() (*tls.Certificate, error) {
	clientCert, err := osutil.GetenvRequiredFromBase64("DOCKER_CLIENTCERT")
	if err != nil {
		return nil, err
	}

	clientCertKey, err := osutil.GetenvRequiredFromBase64("DOCKER_CLIENTCERT_KEY")
	if err != nil {
		return nil, err
	}

	clientKeypair, err := tls.X509KeyPair(clientCert, clientCertKey)
	if err != nil {
		return nil, err
	}

	return &clientKeypair, nil
}
