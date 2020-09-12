package coinbase

import (
	"net/http"
	"net/url"
)

// Authenticator is an interface that objects can implement in order to act as the
// authentication mechanism for RPC requests to Coinbase
type authenticator interface {
	getBaseUrl() *url.URL
	getClient() *http.Client
	authenticate(req *http.Request, method string, requestPath string, body []byte) error
}
