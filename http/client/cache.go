package client

import (
	"net/http"

	"github.com/bartventer/httpcache"
)

type (
	Option = httpcache.Option
)

func CacheTransport(dsn string, ops ...Option) http.RoundTripper {
	return httpcache.NewTransport(dsn, ops...)
}
