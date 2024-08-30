package slog

import (
	"context"
	"log/slog"
	"maps"
	"slices"

	"github.com/gocraft/dbr/v2"
)

type loggerEventReceiver struct {
	ctx context.Context
	sl  *slog.Logger
	lv  slog.Leveler
}

var _ dbr.EventReceiver = (*loggerEventReceiver)(nil)

func New(sl *slog.Logger, lv slog.Leveler) dbr.EventReceiver {
	return &loggerEventReceiver{
		ctx: context.Background(),
		sl:  sl,
		lv:  lv,
	}
}

func (l *loggerEventReceiver) Event(eventName string) {
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke Event", slog.String("event_name", eventName))
}

func (l *loggerEventReceiver) EventKv(eventName string, kvs map[string]string) {
	attrs := kvsToAttrs(kvs, 1)
	attrs = append(attrs, slog.String("event_name", eventName))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErrKv", attrs...)
}

func (l *loggerEventReceiver) EventErr(eventName string, err error) error {
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErr",
		slog.String("event_name", eventName),
		slog.Any("err", err),
	)
	return err
}

func (l *loggerEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	attrs := kvsToAttrs(kvs, 2)
	attrs = append(attrs, slog.String("event_name", eventName), slog.Any("err", err))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErrKv", attrs...)
	return err
}

func (l *loggerEventReceiver) Timing(eventName string, nanoseconds int64) {
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke Timing",
		slog.String("event_name", eventName),
		slog.Int64("nanoseconds", nanoseconds),
	)
}

func (l *loggerEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	attrs := kvsToAttrs(kvs, 2)
	attrs = append(attrs, slog.String("event_name", eventName), slog.Int64("nanoseconds", nanoseconds))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke TimingKv", attrs...)
}

func kvsToAttrs(kvs map[string]string, xlen int) []slog.Attr {
	attrs := make([]slog.Attr, len(kvs)+xlen)
	for i, k := range slices.Collect(maps.Keys(kvs)) {
		attrs[i] = slog.String(k, kvs[k])
	}
	return attrs
}
