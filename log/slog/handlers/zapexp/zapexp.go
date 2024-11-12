package zapexp

import (
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type Handler = zapslog.Handler

type HandlerOption = zapslog.HandlerOption

func New(core zapcore.Core, opts ...HandlerOption) *Handler {
	h := zapslog.NewHandler(core, opts...)
	return h
}
