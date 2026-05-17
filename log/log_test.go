package log

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/phuslu/log"
)

func TestLogger_1(t *testing.T) {
	ll := log.Logger{
		Level:      log.InfoLevel,
		Caller:     1,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
		Writer:     &log.IOWriter{Writer: os.Stdout},
		Context:    log.NewContext(nil).Str("ctx", "some_ctx").Value(),
	}

	ll.Info().Str("key", "val").Msg("日志")
}

func TestLogger_Slog_1(t *testing.T) {
	ll := DefaultLogger

	sll := log.SlogNewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	fmt.Println(sll)

	sl := ll.Slog()
	sl.Info("This is slog", "name", "hello")

	ll.WithLevel(log.InfoLevel).
		Str("key", "val").
		Msg("test")

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

func TestLog_MultiLevelWriter(t *testing.T) {
	ll := log.Logger{Writer: &log.MultiWriter{
		InfoWriter: &log.AsyncWriter{
			Writer: &log.FileWriter{ProcessID: true, Filename: "main.INFO.log", MaxSize: 100 << 20},
		},
		WarnWriter: &log.AsyncWriter{
			Writer: &log.FileWriter{ProcessID: true, Filename: "main.WARN.log", MaxSize: 100 << 20},
		},
		ErrorWriter: &log.AsyncWriter{
			Writer: &log.FileWriter{ProcessID: true, Filename: "main.ERROR.log", MaxSize: 100 << 20},
		},
		ConsoleWriter: &log.ConsoleWriter{ColorOutput: true},
		ConsoleLevel:  log.ErrorLevel,
	},
		Level: InfoLevel,
	}

	sll := ll.Slog()

	sll.Info("info level msg")
	sll.Warn("warning level msg")
	sll.Error("error level msg")
}
