package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

const DefaultTimeout = 5 * time.Second

func HTTP3Transport(tlsConf *tls.Config) http.RoundTripper {
	return HTTP3TransportConf(tlsConf, new(quic.Config))
}

func HTTP3TransportConf(tlsConf *tls.Config, qconf *quic.Config) http.RoundTripper {
	return &http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QUICConfig:      qconf,
	}
}
