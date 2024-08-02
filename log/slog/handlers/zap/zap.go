package zap

import (
	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"
)

type Option = slogzap.Option

func New(o *Option) slog.Handler {
	if o == nil {
		o = &Option{
			Level: slog.LevelInfo,
		}
	}
	return o.NewZapHandler()
}
