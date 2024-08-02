package sentry

import (
	"log/slog"

	"github.com/samber/slog-sentry/v2"
)

type Option = slogsentry.Option

func New(o *Option) slog.Handler {
	if o == nil {
		o = &Option{
			Level: slog.LevelInfo,
		}
	}
	return o.NewSentryHandler()
}
