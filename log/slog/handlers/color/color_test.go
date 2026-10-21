package color

import (
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestColor_1(t *testing.T) {
	crlog := slog.New(New(os.Stderr,
		ShowSource(),
		SetTimeFormat(time.Stamp),
		SetLevel(slog.LevelDebug),
	))

	crlog.Debug("Level Debug with color")
	crlog.Info("Level Info with color")
	crlog.Warn("Level Warn with color")
	crlog.Error("Level Error with color")
}

func TestColor_2(t *testing.T) {
	var ok = false
	slog.Info("debug info", "hhhh", "kkkkk")
	if ok {
		slog.Info("debug info", "hhhh", "used")
	}
	slog.Info("debug info", "aaa", "bbb")
}

func TestSetLevelStr(t *testing.T) {
	crlog := slog.New(New(os.Stderr,
		SetLevelStr(MapOfLevel{
			slog.LevelDebug: "DBG",
			slog.LevelInfo:  "INF",
			slog.LevelWarn:  "WRN",
			slog.LevelError: "ERR",
		}),
	))

	crlog.Info("custom level strings")
}

func TestSetReplaceAttr(t *testing.T) {
	crlog := slog.New(New(os.Stderr,
		SetReplaceAttr(func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "secret" {
				return slog.Attr{}
			}

			return a
		}),
	))

	crlog.Info("redacted", "secret", "hidden", "visible", "shown")
}

func TestDisableColor(t *testing.T) {
	crlog := slog.New(New(os.Stderr, DisableColor()))

	crlog.Info("no color")
}
