package testcerts

import (
	"github.com/madflojo/testcerts"
)

type (
	CertificateAuthority = testcerts.CertificateAuthority
	KeyPair              = testcerts.KeyPair
	KeyPairConfig        = testcerts.KeyPairConfig
)

var (
	GenerateCerts           = testcerts.GenerateCerts
	GenerateCertsToFile     = testcerts.GenerateCertsToFile
	GenerateCertsToTempFile = testcerts.GenerateCertsToTempFile
)
