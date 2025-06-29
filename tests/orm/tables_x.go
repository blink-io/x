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
	Tag sqx.Executor[Tag, TagSetter]
}{
	Tag: sqx.NewExecutor[TAGS, Tag, TagSetter](Tables.Tags),
}
