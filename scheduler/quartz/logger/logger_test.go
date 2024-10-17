package logger

import (
	"log/slog"
	"testing"
)

func TestLogger_1(t *testing.T) {
	ll := New(slog.Default())
	ll.Infof("string: %s", "heison")
}
