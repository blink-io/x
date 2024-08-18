package retryablehttp

import (
	"crypto/tls"

	"github.com/blink-io/x/http/client"
	"github.com/hashicorp/go-retryablehttp"
)

type (
	Client = retryablehttp.Client
)

func HTTP3Client(tlsConf *tls.Config) *retryablehttp.Client {
	cc := retryablehttp.NewClient()
	cc.HTTPClient = client.HTTP3Client(tlsConf)
	//cc.HTTPClient.Transport = client.HTTP3Transport(tlsConf)
	return cc
}
