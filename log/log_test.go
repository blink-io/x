package log

import (
	"fmt"
	"testing"

	"github.com/phuslu/log"
)

func TestLogger_1(t *testing.T) {
	ll := DefaultLogger

	ll.Info().Str("key", "val").Msg("test")
}

func TestLogger_Slog_1(t *testing.T) {
	ll := DefaultLogger

	sl := ll.Slog()

	sl.Info("hello world", "name", "hello")

	ll.SetLevel(log.DebugLevel)

	ll.WithLevel(log.InfoLevel).Str("key", "val").Msg("test")

	xid_ := log.NewXID()
	fmt.Println(xid_.String())
}
