package orm

import (
	"context"
	"time"

	"github.com/blink-io/opt/null"
	sqx "github.com/blink-io/x/sql/builder/sq"
	"github.com/bokwoon95/sq"
)

type mappers struct {
	TAGS sqx.Mapper[TAGS, Tag, TagSetter]
}

var Mappers = mappers{
	TAGS: NewTagMapper(),
}

type TagMapper struct {
	tbl TAGS
}

func NewTagMapper() sqx.Mapper[TAGS, Tag, TagSetter] {
	return TagMapper{
		tbl: sq.New[TAGS](""),
	}
}

func (m TagMapper) columnSetter(c *sq.Column, v TagSetter) {
	tbl := m.tbl
	if iv := v.ID.GetOrZero(); iv > 0 {
		c.SetInt64(tbl.ID, iv)
	}
	v.GUID.IfSet(func(iv string) {
		c.SetString(tbl.GUID, iv)
	})
	v.Name.IfSet(func(iv string) {
		c.SetString(tbl.NAME, iv)
	})
	v.Code.IfSet(func(iv string) {
		c.SetString(tbl.CODE, iv)
	})
	v.CreatedAt.IfSet(func(iv time.Time) {
		c.SetTime(tbl.CREATED_AT, iv)
	})
	v.Description.IfSet(func(t string) {
		c.SetString(tbl.DESCRIPTION, v.Description.GetOrZero())
	})
}

func (m TagMapper) Table() TAGS {
	return m.tbl
}

func (m TagMapper) InsertT(ctx context.Context, vv ...TagSetter) func(*sq.Column) {
	return func(c *sq.Column) {
		for _, v := range vv {
			m.columnSetter(c, v)
		}
	}
}

func (m TagMapper) UpdateT(ctx context.Context, v TagSetter) func(*sq.Column) {
	return func(c *sq.Column) {
		m.columnSetter(c, v)
	}
}

func (m TagMapper) QueryT(ctx context.Context) func(*sq.Row) Tag {
	tbl := m.tbl
	return func(r *sq.Row) Tag {
		v := Tag{
			ID:        r.Int64Field(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			Code:      r.StringField(tbl.CODE),
			Name:      r.StringField(tbl.NAME),
			CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
		desc := r.NullStringField(tbl.DESCRIPTION)
		v.Description = null.FromCond(desc.String, desc.Valid)
		return v
	}
}
