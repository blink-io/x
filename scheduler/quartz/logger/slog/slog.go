package slog

import (
	"context"
	"fmt"
	"log/slog"

	qlogger "github.com/reugn/go-quartz/logger"
)

type logger struct {
	ctx context.Context
	sl  *slog.Logger
}

var _ qlogger.Logger = (*logger)(nil)

func New(sl *slog.Logger) qlogger.Logger {
	return &logger{
		ctx: context.Background(),
		sl:  sl,
	}
}

func (l *logger) Trace(format string, args ...any) {
	l.sl.Log(l.ctx, slog.LevelDebug-4, fmt.Sprintf(format, args...))
}

func (l *logger) Debug(format string, args ...any) {
	l.sl.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Info(format string, args ...any) {
	l.sl.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Warn(format string, args ...any) {
	l.sl.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Error(format string, args ...any) {
	l.sl.Error(fmt.Sprintf(format, args...))
}

func (l *logger) Enabled(level qlogger.Level) bool {
	return l.sl.Enabled(l.ctx, toSlogLevel(level))
}

func toSlogLevel(l qlogger.Level) slog.Level {
	switch l {
	case qlogger.LevelTrace:
		return slog.LevelDebug - 4
	case qlogger.LevelDebug:
		return slog.LevelDebug
	case qlogger.LevelInfo:
		return slog.LevelInfo
	case qlogger.LevelWarn:
		return slog.LevelWarn
	case qlogger.LevelError:
		return slog.LevelError
	default:
		return slog.LevelError + 4
	}
}
