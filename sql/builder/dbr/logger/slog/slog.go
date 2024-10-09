package slog

import (
	"context"
	"log/slog"
	"maps"
	"slices"
	"time"

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
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke Event", eventNameToAttr(eventName))
}

func (l *loggerEventReceiver) EventKv(eventName string, kvs map[string]string) {
	attrs := kvsToAttrs(kvs, 1)
	attrs = append(attrs, eventNameToAttr(eventName))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErrKv", attrs...)
}

func (l *loggerEventReceiver) EventErr(eventName string, err error) error {
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErr",
		eventNameToAttr(eventName),
		errToAttr(err),
	)
	return err
}

func (l *loggerEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	attrs := kvsToAttrs(kvs, 2)
	attrs = append(attrs, eventNameToAttr(eventName), errToAttr(err))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke EventErrKv", attrs...)
	return err
}

func (l *loggerEventReceiver) Timing(eventName string, nanoseconds int64) {
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke Timing",
		eventNameToAttr(eventName),
		nsToAttr(nanoseconds),
	)
}

func (l *loggerEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	attrs := kvsToAttrs(kvs, 2)
	attrs = append(attrs, eventNameToAttr(eventName), nsToAttr(nanoseconds))
	l.sl.LogAttrs(l.ctx, l.lv.Level(), "Invoke TimingKv", attrs...)
}

func errToAttr(err error) slog.Attr {
	return slog.Any("err", err)
}

func nsToAttr(ns int64) slog.Attr {
	return slog.Int64("ms", time.Duration(ns).Milliseconds())
}

func eventNameToAttr(eventName string) slog.Attr {
	return slog.String("event_name", eventName)
}

func kvsToAttrs(kvs map[string]string, xlen int) []slog.Attr {
	nlen := len(kvs) + xlen
	attrs := make([]slog.Attr, nlen, nlen)
	for i, k := range slices.Collect(maps.Keys(kvs)) {
		attrs[i] = slog.String(k, kvs[k])
	}
	return attrs
}
