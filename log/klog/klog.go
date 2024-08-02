package klog

import (
	"os"

	"github.com/go-kit/log"
)

type Logger = log.Logger

func Noop() Logger {
	return log.NewNopLogger()
}

func JSON() Logger {
	return log.NewJSONLogger(os.Stdout)
}
