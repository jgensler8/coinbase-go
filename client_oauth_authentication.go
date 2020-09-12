package coinbase

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

// ClientOAuthAuthentication Struct implements the Authentication interface
// and takes care of authenticating OAuth RPC requests on behalf of a client
// (i.e GetBalance())
type clientOAuthAuthentication struct {
	Tokens  *oauthTokens
	BaseUrl *url.URL
	Client  http.Client
}

// ClientOAuth instantiates ClientOAuthAuthentication with the client OAuth tokens
func clientOAuth(tokens *oauthTokens) *clientOAuthAuthentication {
	a := clientOAuthAuthentication{
		Tokens:  tokens,
		BaseUrl: ProURL(false, URLTypeWebsite),
		Client: http.Client{
			Transport: &http.Transport{
				Dial: dialTimeout,
			},
		},
	}
	return &a
}

// Client OAuth authentication requires us to attach an unexpired OAuth token to
// the request header
func (a clientOAuthAuthentication) authenticate(req *http.Request, method string, requestPath string, body []byte) error {
	// Ensure tokens havent expired
	if time.Now().UTC().Unix() > a.Tokens.ExpireTime {
		return errors.New("The OAuth tokens are expired. Use refreshTokens to refresh them")
	}
	req.Header.Set("Authorization", "Bearer "+a.Tokens.AccessToken)
	return nil
}

func (a clientOAuthAuthentication) getBaseUrl() *url.URL {
	return copyURL(a.BaseUrl)
}

func (a clientOAuthAuthentication) getClient() *http.Client {
	return &a.Client
}
