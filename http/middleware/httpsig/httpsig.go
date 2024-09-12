package httpsig

import "github.com/42wim/httpsig"

var (
	NewVerifier         = httpsig.NewVerifier
	NewResponseVerifier = httpsig.NewResponseVerifier
	NewSigner           = httpsig.NewSigner
	NewSSHSigner        = httpsig.NewSSHSigner
)
