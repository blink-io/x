package orm

import (
	"github.com/blink-io/sqx"
)

const (
	EnumEnumsStatusActive EnumEnumsStatus = "active"
)

func init() {
	EnumEnumsStatusValues = append(EnumEnumsStatusValues, string(EnumEnumsStatusActive))
}

var Executors = struct {
	Tag   sqx.Executor[Tag, TagSetter]
	Log   sqx.Executor[Log, LogSetter]
	Array sqx.Executor[Array, ArraySetter]
}{
	Tag:   sqx.NewExecutor[TAGS, Tag, TagSetter](Tables.Tags),
	Log:   sqx.NewExecutor[LOGS, Log, LogSetter](Tables.Logs),
	Array: sqx.NewExecutor[ARRAYS, Array, ArraySetter](Tables.Arrays),
}
