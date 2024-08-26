package passwap

import "github.com/zitadel/passwap"

var (
	ErrPasswordMismatch = passwap.ErrPasswordMismatch
	ErrPasswordNoChange = passwap.ErrPasswordNoChange
	ErrNoVerifier       = passwap.ErrNoVerifier

	NewSwapper = passwap.NewSwapper
)

type (
	Hasher     = passwap.Hasher
	SkipErrors = passwap.SkipErrors
	Swapper    = passwap.Swapper
)
