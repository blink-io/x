package std

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

const (
	prefixDebug = "DEBUG :"
	prefixInfo  = "INFO :"
	prefixWarn  = "WARN :"
	prefixError = "ERROR :"
)

func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			msg = prefixDebug + msg
		case logging.LevelInfo:
			msg = prefixInfo + msg
		case logging.LevelWarn:
			msg = prefixWarn + msg
		case logging.LevelError:
			msg = prefixError + msg
		default:
			// A config mistake must not crash the request path; log it
			// through the same channel with a marker prefix.
			msg = prefixWarn + "UNKNOWN-LEVEL " + msg
		}
		l.Println(append([]any{"msg", msg}, fields...))
	})
}
