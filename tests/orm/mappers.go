package orm

import (
	"context"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/sq"
	"github.com/blink-io/sqx"
)

type mappers struct {
	TAGS sqx.Mapper[TAGS, Tag, TagSetter]
}

var Mappers = mappers{
	TAGS: NewTagMapper(),
}
