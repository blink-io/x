package log

import (
	"fmt"
	"os"
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
	sl.Info("This is slog", "name", "hello")

	ll.WithLevel(log.InfoLevel).Str("key", "val").Msg("test")

	xid_ := log.NewXID()
	fmt.Println(xid_.String())
}

func TestLog_FileWriter(t *testing.T) {
	fw := &FileWriter{
		//LocalTime: true,
		//HostName:  true,
		//ProcessID: true,
		Filename: "test.log",
	}

	mw := &MultiIOWriter{
		IOWriter{Writer: os.Stderr},
		fw,
	}

	ll := &Logger{
		Level:  InfoLevel,
		Writer: mw,
	}
	ll.Info().Msg("test")

}
