package logrus

import (
	"log/slog"

	sloglogrus "github.com/samber/slog-logrus/v2"
)

type Option = sloglogrus.Option

func New(o *Option) slog.Handler {
	if o == nil {
		o = &Option{
			Level: slog.LevelInfo,
		}
	}
	return o.NewLogrusHandler()
}
