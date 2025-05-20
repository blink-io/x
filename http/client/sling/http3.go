package sling

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/blink-io/x/http/client"
	"github.com/dghubble/sling"
	"github.com/quic-go/quic-go"
)

func HTTP3Client(tlsConf *tls.Config) *sling.Sling {
	return HTTP3ClientWithConf(tlsConf, new(quic.Config))
}

func HTTP3ClientWithConf(tlsConf *tls.Config, qconf *quic.Config) *sling.Sling {
	cc := sling.New().Client(
		&http.Client{
			Timeout:   5 * time.Second,
			Transport: client.HTTP3TransportWithConf(tlsConf, qconf),
		},
	)
	return cc
}
