package zapexp

import (
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type Handler = zapslog.Handler

type Options = zapslog.HandlerOptions

func New(core zapcore.Core, opts *Options) *Handler {
	h := zapslog.NewHandler(core, opts)
	return h
}
