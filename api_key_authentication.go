package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ApiKeyAuthentication Struct implements the Authentication interface and takes
// care of authenticating RPC requests for clients with a Key & Secret pair
type apiKeyAuthentication struct {
	Key        string
	Secret     string
	Passphrase string
	BaseUrl    *url.URL
	Client     http.Client
}

// ApiKeyAuth instantiates ApiKeyAuthentication with the API key & secret
func apiKeyAuth(key string, secret string, passphrase string) *apiKeyAuthentication {
	a := apiKeyAuthentication{
		Key:        key,
		Secret:     secret,
		Passphrase: passphrase,
		BaseUrl:    ProURL(false, URLTypeREST),
		Client: http.Client{
			Transport: &http.Transport{
				Dial: dialTimeout,
			},
		},
	}
	return &a
}

// API Key + Secret authentication requires a request header of the HMAC SHA-256
// signature of the "message" as well as an incrementing nonce and the API key
// params:
// req: request to set headers on
// method: POST
// requestPathAndBody: `/orders{"key":"value"}`
func (a apiKeyAuthentication) authenticate(req *http.Request, method string, requestPath string, body []byte) error {

	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	// create the prehash string by concatenating required parts
	what := timestamp + method + requestPath
	if len(body) > 0 {
		what += string(body)
	}
	fmt.Printf("method:%v:requestPath:%v:body:%v:what:%v:\n", method, requestPath, string(body), what)

	// decode the base64 secret
	hmacKey, err := base64.StdEncoding.DecodeString(a.Secret)
	if err != nil {
		return err
	}

	// create a sha256 hmac with the secret
	h := hmac.New(sha256.New, hmacKey)
	h.Write([]byte(what))

	// sign the require message with the hmac
	// and finally base64 encode the result
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// The api key as a string.
	req.Header.Set("CB-ACCESS-KEY", a.Key)
	// The base64-encoded signature (see Signing a Message).
	req.Header.Set("CB-ACCESS-SIGN", signature)
	// A timestamp for your request.
	req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)
	// The passphrase you specified when creating the API key.
	req.Header.Set("CB-ACCESS-PASSPHRASE", a.Passphrase)

	return nil
}

func (a apiKeyAuthentication) getBaseUrl() *url.URL {
	return copyURL(a.BaseUrl)
}

func (a apiKeyAuthentication) getClient() *http.Client {
	return &a.Client
}
