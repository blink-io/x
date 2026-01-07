package slog

import (
	"context"
	"fmt"
	"log/slog"

	klog "github.com/go-kratos/kratos/v2/log"
)

type logger struct {
	log *slog.Logger
}

var _ klog.Logger = (*logger)(nil)

func NewLogger(log *slog.Logger) klog.Logger {
	return newLogger(log)
}

func newLogger(log *slog.Logger) *logger {
	return &logger{
		log: log,
	}
}

func (l *logger) Log(level klog.Level, keyvals ...any) error {
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("keyvals must appear in pairs: ", keyvals))
		return nil
	}

	args := make([]slog.Attr, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		args = append(args, slog.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	lv := slog.LevelInfo
	switch level {
	case klog.LevelDebug:
		lv = slog.LevelDebug
	case klog.LevelInfo:
		lv = slog.LevelInfo
	case klog.LevelWarn:
		lv = slog.LevelWarn
	case klog.LevelError:
		lv = slog.LevelError
	case klog.LevelFatal:
		lv = slog.LevelError + 4
	}

	var msg string
	if level == klog.LevelFatal {
		msg = "[Fatal]"
		args = append(args, slog.String("LOG_LEVEL", klog.LevelFatal.String()))
	}

	l.log.LogAttrs(context.Background(), lv, msg, args...)

	return nil
}
