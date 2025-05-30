package color

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"gitlab.com/greyxor/slogor"
)

func TestColor_1(t *testing.T) {
	crlog := slog.New(slogor.NewHandler(os.Stderr,
		slogor.ShowSource(),
		slogor.SetTimeFormat(time.Stamp),
		slogor.SetLevel(slog.LevelDebug),
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
		slog.Info("debug info", "hhhh")
	}
	slog.Info("debug info", "aaa", "bbb")
}
