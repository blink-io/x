package go_password

import "github.com/sethvargo/go-password/password"

var (
	ErrExceedsTotalLength      = password.ErrExceedsTotalLength
	ErrLettersExceedsAvailable = password.ErrLettersExceedsAvailable
	ErrDigitsExceedsAvailable  = password.ErrDigitsExceedsAvailable
	ErrSymbolsExceedsAvailable = password.ErrSymbolsExceedsAvailable

	NewGenerator = password.NewGenerator
	Generate     = password.Generate
)

type (
	Generator         = password.Generator
	GeneratorInput    = password.GeneratorInput
	PasswordGenerator = password.PasswordGenerator
)
