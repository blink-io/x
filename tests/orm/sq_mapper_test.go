package orm

import (
	"github.com/aarondl/opt/omitnull"
	"github.com/bokwoon95/sq"
)

type Mapper[T sq.Table, M any] interface {
	Table() T
	InsertMapper(...M) func(*sq.Column)
	QueryMapper() func(*sq.Row) M
}

type TagMapper struct {
	tbl TAGS
}

func NewTagMapper() Mapper[TAGS, Tag] {
	return TagMapper{
		tbl: sq.New[TAGS]("tags"),
	}
}

func (m TagMapper) Table() TAGS {
	return m.tbl
}

func (m TagMapper) InsertMapper(vv ...Tag) func(*sq.Column) {
	tbl := m.tbl
	return func(c *sq.Column) {
		for _, v := range vv {
			if v.ID > 0 {
				c.SetInt(tbl.ID, v.ID)
			}
			c.SetString(tbl.GUID, v.GUID)
			c.SetString(tbl.NAME, v.Name)
			c.SetString(tbl.CODE, v.Code)
			c.SetTime(tbl.CREATED_AT, v.CreatedAt)
			c.SetString(tbl.DESCRIPTION, v.Description.GetOrZero())
		}
	}
}

func (m TagMapper) QueryMapper() func(*sq.Row) Tag {
	tbl := m.tbl
	return func(r *sq.Row) Tag {
		u := Tag{
			ID:          r.IntField(tbl.ID),
			GUID:        r.StringField(tbl.GUID),
			Code:        r.StringField(tbl.CODE),
			Name:        r.StringField(tbl.NAME),
			Description: omitnull.From(r.NullStringField(tbl.DESCRIPTION).String),
		}
		return u
	}
}
