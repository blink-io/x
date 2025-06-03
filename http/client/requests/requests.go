package requests

import (
	"crypto/tls"

	"github.com/blink-io/x/http/client"

	"github.com/carlmjohnson/requests"
	"github.com/quic-go/quic-go"
)

type Builder = requests.Builder

func HTTP3(tlsConf *tls.Config) *Builder {
	return HTTP3Conf(tlsConf, new(quic.Config))
}

func HTTP3Conf(tlsConf *tls.Config, qconf *quic.Config) *Builder {
	cc := requests.New().
		Transport(client.HTTP3TransportWithConf(tlsConf, qconf))
	return cc
}
