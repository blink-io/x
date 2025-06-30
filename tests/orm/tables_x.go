package orm

import (
	"github.com/blink-io/sqx"
)

const (
	EnumEnumsStatusActive  EnumEnumsStatus = "active"
	EnumEnumsStatusBlocked EnumEnumsStatus = "blocked"
)

func init() {
	EnumEnumsStatusValues = append(EnumEnumsStatusValues, string(EnumEnumsStatusActive))
	EnumEnumsStatusValues = append(EnumEnumsStatusValues, string(EnumEnumsStatusBlocked))
}

var Executors = struct {
	Tag   sqx.Executor[TAGS, Tag, TagSetter]
	Log   sqx.Executor[LOGS, Log, LogSetter]
	Array sqx.Executor[ARRAYS, Array, ArraySetter]
	Enum  sqx.Executor[ENUMS, Enum, EnumSetter]
	MKey  sqx.Executor[MKEYS, Mkey, MkeySetter]
}{
	Tag:   sqx.NewExecutor[TAGS, Tag, TagSetter](Tables.Tags),
	Log:   sqx.NewExecutor[LOGS, Log, LogSetter](Tables.Logs),
	Array: sqx.NewExecutor[ARRAYS, Array, ArraySetter](Tables.Arrays),
	Enum:  sqx.NewExecutor[ENUMS, Enum, EnumSetter](Tables.Enums),
	MKey:  sqx.NewExecutor[MKEYS, Mkey, MkeySetter](Tables.Mkeys),
}
