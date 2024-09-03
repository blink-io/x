package orm

import (
	"github.com/aarondl/opt/omitnull"
	"github.com/bokwoon95/sq"
)

type Mapper[T sq.Table, M any] interface {
	Table() T
	InsertColumns(...M) func(*sq.Column)
	RowMapper() func(*sq.Row) M
}

type UserMappers struct {
	tbl USERS
}

func NewUserMappers(tbl USERS) UserMappers {
	return UserMappers{tbl: tbl}
}

func (m UserMappers) Table() USERS {
	return m.tbl
}

func (m UserMappers) InsertColumns(r *User) func(*sq.Column) {
	tbl := m.tbl
	return func(col *sq.Column) {
		col.SetString(tbl.GUID, r.GUID)
		col.SetString(tbl.USERNAME, r.Username)
		col.SetFloat64(tbl.SCORE, r.Score)
		col.SetTime(tbl.CREATED_AT, r.CreatedAt)
		col.SetTime(tbl.UPDATED_AT, r.UpdatedAt)
	}
}

func (m UserMappers) RowMapper() func(*sq.Row) *User {
	tbl := m.tbl
	return func(r *sq.Row) *User {
		u := &User{
			ID:        r.IntField(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			Username:  r.StringField(tbl.USERNAME),
			Score:     r.Float64Field(tbl.SCORE),
			CreatedAt: r.TimeField(tbl.CREATED_AT),
			UpdatedAt: r.TimeField(tbl.UPDATED_AT),
		}
		return u
	}
}

type TagMappers struct {
	tbl TAGS
}

func NewTagMappers(tbl TAGS) Mapper[TAGS, Tag] {
	return TagMappers{tbl: tbl}
}

func (m TagMappers) Table() TAGS {
	return m.tbl
}

func (m TagMappers) InsertColumns(vv ...Tag) func(*sq.Column) {
	tbl := m.tbl
	return func(col *sq.Column) {
		for _, r := range vv {
			col.SetString(tbl.GUID, r.GUID)
			col.SetString(tbl.NAME, r.Name)
			col.SetString(tbl.CODE, r.Code)
			col.SetString(tbl.DESCRIPTION, r.Description.GetOrZero())
		}
	}
}

func (m TagMappers) RowMapper() func(*sq.Row) Tag {
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
