package log

import "testing"

func TestLogger_1(t *testing.T) {
	ll := DefaultLogger

	ll.Info().Str("key", "val").Msg("test")
}

func TestLogger_Slog_1(t *testing.T) {
	ll := DefaultLogger

	sl := ll.Slog()

	sl.Info("hello world", "name", "hello")
}
