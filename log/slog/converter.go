package slog

import (
	"log/slog"

	slogcommon "github.com/samber/slog-common"
)

var (
	SourceKey = "source"
	ErrorKeys = []string{"error", "err"}
)

type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) []any

func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) []any {
	// aggregate all attributes
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	// developer formatters
	attrs = slogcommon.ReplaceError(attrs, ErrorKeys...)
	if addSource {
		attrs = append(attrs, slogcommon.Source(SourceKey, record))
	}
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)

	// handler formatter
	return attrsToPairs(attrs)
}

// attrsToPairs flattens attrs into an ordered key/value slice.
//
// It follows slogcommon.AttrsToMap's merge rules: a repeated key keeps the
// last scalar value, and repeated groups are merged attribute by attribute.
// Unlike AttrsToMap it keeps the order in which the keys first appear.
// AttrsToMap returns a map, and Go randomises map iteration order, so
// flattening that map emitted the same record with its fields in a different
// order on every call.
//
// Group values stay nested maps, exactly as AttrsToMap produces them, so the
// output shape is unchanged.
func attrsToPairs(attrs []slog.Attr) []any {
	keys := make([]string, 0, len(attrs))
	values := make([]slog.Value, 0, len(attrs))
	positions := make(map[string]int, len(attrs))

	for _, attr := range attrs {
		if i, ok := positions[attr.Key]; ok {
			values[i] = mergeValues(values[i], attr.Value)
			continue
		}

		positions[attr.Key] = len(keys)
		keys = append(keys, attr.Key)
		values = append(values, attr.Value)
	}

	output := make([]any, 0, len(keys)*2)
	for i, key := range keys {
		v := values[i]
		if v.Kind() == slog.KindGroup {
			output = append(output, key, slogcommon.AttrsToMap(v.Group()...))
			continue
		}

		output = append(output, key, v.Any())
	}

	return output
}

// mergeValues merges two values that share a key, following the rule
// slogcommon.AttrsToMap applies when a key appears more than once.
//
// Two groups are merged attribute by attribute, anything else keeps the last
// value seen.
func mergeValues(a, b slog.Value) slog.Value {
	if a.Kind() != slog.KindGroup || b.Kind() != slog.KindGroup {
		return b
	}

	group := make([]slog.Attr, 0, len(a.Group())+len(b.Group()))
	group = append(group, a.Group()...)
	group = append(group, b.Group()...)

	return slog.GroupValue(group...)
}
