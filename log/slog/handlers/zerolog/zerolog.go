package zerolog

import (
	"log/slog"

	"github.com/samber/slog-zerolog/v2"
)

type Option = slogzerolog.Option

func New(o *Option) slog.Handler {
	if o == nil {
		o = &Option{
			Level: slog.LevelInfo,
		}
	}
	return o.NewZerologHandler()
}
