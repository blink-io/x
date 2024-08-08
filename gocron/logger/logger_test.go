package logger

import (
	"log/slog"
	"testing"
)

func TestLogger_1(t *testing.T) {
	ll := New(slog.Default())

	ll.Info("Hello", "abc", 1111, 222, true, "ok", 3.14)
}
