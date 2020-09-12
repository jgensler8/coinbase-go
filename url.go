package coinbase

import (
	"net/url"
)

type URLType int

const (
	URLTypeWebsite   URLType = iota
	URLTypeOAuth     URLType = iota
	URLTypeREST      URLType = iota
	URLTypeWebsocket URLType = iota
	URLTypeFIX       URLType = iota
)

func mustParseURL(u string) *url.URL {
	p, _ := url.Parse(u)
	return p
}

// ProURL returns the URL for a given service.
// returns nil if no URLType matched
// setting sandbox to `true` will use Coinbase's sandbox environment
// URLType should be one specified in this package (Website, REST, ...)
func ProURL(sandbox bool, urlType URLType) *url.URL {
	switch urlType {
	case URLTypeWebsite:
		if sandbox {
			return mustParseURL("https://public.sandbox.pro.coinbase.com")
		}
		// TODO(jeffg): figure out if this is the case
		return mustParseURL("https://pro.coinbase.com")
	case URLTypeOAuth:
		if sandbox {
			return nil
		}
		return mustParseURL("https://coinbase.com")
	case URLTypeREST:
		if sandbox {
			return mustParseURL("https://api-public.sandbox.pro.coinbase.com")
		}
		return mustParseURL("https://api.pro.coinbase.com")
	case URLTypeWebsocket:
		if sandbox {
			return mustParseURL("wss://ws-feed-public.sandbox.pro.coinbase.com")
		}
		return mustParseURL("wss://ws-feed.pro.coinbase.com")
	case URLTypeFIX:
		if sandbox {
			return mustParseURL("tcp+ssl://fix-public.sandbox.pro.coinbase.com:4198")
		}
		return mustParseURL("tcp+ssl://fix.pro.coinbase.com:4198")
	}
	return nil
}
