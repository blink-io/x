package resty

import (
	"crypto/tls"

	"github.com/blink-io/x/http/client"

	"github.com/go-resty/resty/v2"
	"github.com/quic-go/quic-go"
)

func HTTP3Client(tlsConf *tls.Config) *resty.Client {
	return HTTP3ClientWithConf(tlsConf, new(quic.Config))
}

func HTTP3ClientWithConf(tlsConf *tls.Config, qconf *quic.Config) *resty.Client {
	return resty.New().
		SetTransport(client.HTTP3TransportWithConf(tlsConf, qconf))
}
