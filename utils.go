package coinbase

import (
	"net"
	"net/url"
	"time"
)

// dialTimeout is used to enforce a timeout for all http requests.
func dialTimeout(network, addr string) (net.Conn, error) {
	var timeout = time.Duration(2 * time.Second) //how long to wait when trying to connect to the coinbase
	return net.DialTimeout(network, addr, timeout)
}

func copyURL(u *url.URL) *url.URL {
	return mustParseURL(u.String())
}
