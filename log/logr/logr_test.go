package logr

import (
	"log/slog"
	"os"
	"testing"
)

func TestLogr_Slog_1(t *testing.T) {
	var hdlr slog.Handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	var sl = FromSlogHandler(hdlr)

	sl.Info("Hello")

}
