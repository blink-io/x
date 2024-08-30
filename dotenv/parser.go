package dotenv

import (
	goenv "github.com/Netflix/go-env"
	"github.com/caarlos0/env/v10"
)

var (
	Parse                     = env.Parse
	ParseWithOptions          = env.ParseWithOptions
	GetFieldParams            = env.GetFieldParams
	GetFieldParamsWithOptions = env.GetFieldParamsWithOptions

	Marshal              = goenv.Marshal
	Unmarshal            = goenv.Unmarshal
	UnmarshalFromEnviron = goenv.UnmarshalFromEnviron
	EnvSetToEnviron      = goenv.EnvSetToEnviron
)
