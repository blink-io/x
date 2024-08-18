package logr

import (
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/go-logr/stdr"
)

func TestLogr_Slog_1(t *testing.T) {
	var hdlr slog.Handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	var sl = FromSlogHandler(hdlr)

	sl.Info("Hello")

}

func TestLogr_Slog_2(t *testing.T) {
	ll := stdr.New(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
	sl := slog.New(ToSlogHandler(ll))

	sl.Info("Hello")
}
