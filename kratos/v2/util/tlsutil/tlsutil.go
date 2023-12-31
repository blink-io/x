package tlsutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
)

func GenerateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"kratos-quic-server"},
	}
}

func MustInsecureTLSConfig() *tls.Config {
	tlsConf, err := InsecureTLSConfig()
	if err != nil {
		panic(err)
	}
	return tlsConf
}

func InsecureTLSConfig() (*tls.Config, error) {
	if pool, err := x509.SystemCertPool(); err != nil {
		return nil, err
	} else {
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            pool,
		}
		return tlsConf, nil
	}
}
