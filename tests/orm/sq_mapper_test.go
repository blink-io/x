package orm

import "github.com/blink-io/sqx"

type mappers struct {
	TAGS sqx.Mapper[TAGS, Tag, TagSetter]
}

var Mappers = mappers{
	TAGS: NewTagMapper(),
}
