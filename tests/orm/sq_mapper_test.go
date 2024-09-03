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

type UserMapper struct {
	tbl USERS
}

var _ Mapper[USERS, User] = (*UserMapper)(nil)

func NewUserMapper() Mapper[USERS, User] {
	return UserMapper{
		tbl: sq.New[USERS]("users"),
	}
}

func (m UserMapper) Table() USERS {
	return m.tbl
}

func (m UserMapper) InsertMapper(vv ...User) func(*sq.Column) {
	tbl := m.tbl
	return func(col *sq.Column) {
		for _, v := range vv {
			col.SetString(tbl.GUID, v.GUID)
			col.SetString(tbl.USERNAME, v.Username)
			col.SetFloat64(tbl.SCORE, v.Score)
			col.SetTime(tbl.CREATED_AT, v.CreatedAt)
			col.SetTime(tbl.UPDATED_AT, v.UpdatedAt)
		}
	}
}

func (m UserMapper) QueryMapper() func(*sq.Row) User {
	tbl := m.tbl
	return func(r *sq.Row) User {
		u := User{
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
	return func(col *sq.Column) {
		for _, v := range vv {
			col.SetString(tbl.GUID, v.GUID)
			col.SetString(tbl.NAME, v.Name)
			col.SetString(tbl.CODE, v.Code)
			col.SetString(tbl.DESCRIPTION, v.Description.GetOrZero())
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
