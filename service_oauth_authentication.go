package coinbase

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ServiceOAuthAuthentication Struct implements the Authentication interface
// and takes care of authenticating OAuth RPC requests on behalf of the service
// (i.e GetTokens())
type serviceOAuthAuthentication struct {
	BaseUrl *url.URL
	Client  http.Client
}

// ServiceOAuth instantiates ServiceOAuthAuthentication with the coinbase certificate file
func serviceOAuth(certFilePath string) (*serviceOAuthAuthentication, error) {
	// First we read the cert
	certs := x509.NewCertPool()
	pemData, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	mTLSConfig := &tls.Config{
		RootCAs: certs, //Add the cert as a TLS config
	}
	a := serviceOAuthAuthentication{
		BaseUrl: ProURL(false, URLTypeOAuth),
		Client: http.Client{
			Transport: &http.Transport{
				Dial:            dialTimeout,
				TLSClientConfig: mTLSConfig,
			},
		},
	}
	return &a, nil
}

// Service OAuth authentication requires no additional headers to be sent. The
// Coinbase Public Certificate is set as a TLS config in the http.Client
func (a serviceOAuthAuthentication) authenticate(req *http.Request, method string, requestPath string, body []byte) error {
	return nil // No additional headers needed for service OAuth requests
}

func (a serviceOAuthAuthentication) getBaseUrl() *url.URL {
	return copyURL(a.BaseUrl)
}

func (a serviceOAuthAuthentication) getClient() *http.Client {
	return &a.Client
}
