package logger

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-co-op/gocron/v2"
)

type logger struct {
	sl *slog.Logger
}

func New(sl *slog.Logger) gocron.Logger {
	return &logger{sl: sl}
}

func (l *logger) Debug(msg string, args ...any) {
	v := fmt.Sprintf("DEBUG: %s%s\n", msg, logFormatArgs(args...))
	l.sl.Debug(v)
}

func (l *logger) Error(msg string, args ...any) {
	v := fmt.Sprintf("ERROR: %s%s\n", msg, logFormatArgs(args...))
	l.sl.Error(v)
}

func (l *logger) Info(msg string, args ...any) {
	v := fmt.Sprintf("INFO: %s%s\n", msg, logFormatArgs(args...))
	l.sl.Info(v)
}

func (l *logger) Warn(msg string, args ...any) {
	v := fmt.Sprintf("WARN: %s%s\n", msg, logFormatArgs(args...))
	l.sl.Warn(v)
}

func logFormatArgs(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	if len(args)%2 != 0 {
		return ", " + fmt.Sprint(args...)
	}
	var pairs []string
	for i := 0; i < len(args); i += 2 {
		pairs = append(pairs, fmt.Sprintf("%s=%v", args[i], args[i+1]))
	}
	return ", " + strings.Join(pairs, ", ")
}
