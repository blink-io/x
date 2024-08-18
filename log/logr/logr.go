package logr

import "github.com/go-logr/logr"

var (
	FromContextAsSlogLogger = logr.FromContextAsSlogLogger

	NewContext = logr.NewContext

	NewContextWithSlogLogger = logr.NewContextWithSlogLogger

	FromSlogHandler = logr.FromSlogHandler

	ToSlogHandler = logr.ToSlogHandler

	New = logr.New
)

type (
	CallDepthLogSink       = logr.CallDepthLogSink
	CallStackHelperLogSink = logr.CallStackHelperLogSink
	LogSink                = logr.LogSink
	Logger                 = logr.Logger
	Marshaler              = logr.Marshaler
)
