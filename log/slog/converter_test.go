package slog

import (
	"errors"
	"log/slog"
	"reflect"
	"testing"
	"time"

	slogcommon "github.com/samber/slog-common"
)

// legacyConverter is the original implementation, kept here as the reference
// for the differential test below.
func legacyConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) []any {
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	attrs = slogcommon.ReplaceError(attrs, ErrorKeys...)
	if addSource {
		attrs = append(attrs, slogcommon.Source(SourceKey, record))
	}
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)

	fields := slogcommon.AttrsToMap(attrs...)

	output := make([]any, 0)
	for k, v := range fields {
		output = append(output, k, v)
	}

	return output
}

// pairsToMap turns the key/value slice back into a map so it can be compared
// with the map the legacy implementation produced.
func pairsToMap(pairs []any) map[string]any {
	out := make(map[string]any, len(pairs)/2)
	for i := 0; i+1 < len(pairs); i += 2 {
		key, _ := pairs[i].(string)
		out[key] = pairs[i+1]
	}

	return out
}

func TestDefaultConverterMatchesLegacy(t *testing.T) {
	err := errors.New("boom")

	record := slog.NewRecord(time.Time{}.Add(0), slog.LevelInfo, "hello", 0)
	record.AddAttrs(
		slog.String("a", "1"),
		slog.Int("b", 2),
		slog.String("c", "3"),
	)

	loggerAttrs := []slog.Attr{
		slog.String("logger", "yes"),
		slog.Any("err", err),
	}

	groups := []string{"g1", "g2"}

	replaceAttr := func(gs []string, a slog.Attr) slog.Attr {
		if a.Key == "drop" {
			return slog.Attr{}
		}

		return a
	}

	record.AddAttrs(slog.String("drop", "me"), slog.Group("grp", slog.String("x", "y")))

	tests := []struct {
		name        string
		addSource   bool
		replaceAttr func(groups []string, a slog.Attr) slog.Attr
		loggerAttr  []slog.Attr
		groups      []string
	}{
		{name: "plain"},
		{name: "with source", addSource: true},
		{name: "with logger attrs", loggerAttr: loggerAttrs},
		{name: "with groups", groups: groups},
		{name: "with replace attr", replaceAttr: replaceAttr},
		{
			name:        "everything",
			addSource:   true,
			replaceAttr: replaceAttr,
			loggerAttr:  loggerAttrs,
			groups:      groups,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultConverter(tt.addSource, tt.replaceAttr, tt.loggerAttr, tt.groups, &record)
			want := legacyConverter(tt.addSource, tt.replaceAttr, tt.loggerAttr, tt.groups, &record)

			if gotMap, wantMap := pairsToMap(got), pairsToMap(want); !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("DefaultConverter() = %v, want %v", gotMap, wantMap)
			}
		})
	}
}

func TestDefaultConverterOrderIsStable(t *testing.T) {
	record := slog.NewRecord(time.Time{}.Add(0), slog.LevelInfo, "hello", 0)
	record.AddAttrs(
		slog.String("a", "1"),
		slog.String("b", "2"),
		slog.String("c", "3"),
	)

	want := []any{"a", "1", "b", "2", "c", "3"}

	for i := 0; i < 100; i++ {
		got := DefaultConverter(false, nil, nil, nil, &record)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("DefaultConverter() = %v, want %v", got, want)
		}
	}
}

func TestDefaultConverterDuplicateKeys(t *testing.T) {
	record := slog.NewRecord(time.Time{}.Add(0), slog.LevelInfo, "hello", 0)
	record.AddAttrs(
		slog.String("a", "1"),
		slog.String("a", "2"),
		slog.Group("g", slog.String("x", "1")),
		slog.Group("g", slog.String("y", "2")),
	)

	got := DefaultConverter(false, nil, nil, nil, &record)

	want := []any{
		"a", "2",
		"g", map[string]any{"x": "1", "y": "2"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("DefaultConverter() = %v, want %v", got, want)
	}
}

func BenchmarkDefaultConverter(b *testing.B) {
	record := slog.NewRecord(time.Time{}.Add(0), slog.LevelInfo, "hello", 0)
	record.AddAttrs(
		slog.String("a", "1"),
		slog.Int("b", 2),
		slog.String("c", "3"),
		slog.String("d", "4"),
	)

	b.Run("without source", func(b *testing.B) {
		for b.Loop() {
			_ = DefaultConverter(false, nil, nil, nil, &record)
		}
	})

	b.Run("with source", func(b *testing.B) {
		for b.Loop() {
			_ = DefaultConverter(true, nil, nil, nil, &record)
		}
	})
}
