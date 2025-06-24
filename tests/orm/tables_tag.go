package orm

import (
	"context"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/sq"
	"github.com/blink-io/sqx"
)

type Tag struct {
	ID          int64            `db:"id"`
	GUID        string           `db:"guid"`
	Code        string           `db:"code"`
	Name        string           `db:"name"`
	CreatedAt   time.Time        `db:"created_at"`
	Description null.Val[string] `db:"description"`
}

func (m Tag) Setter() TagSetter {
	ss := TagSetter{
		GUID:        omit.From(m.GUID),
		Code:        omit.From(m.Code),
		Name:        omit.From(m.Name),
		CreatedAt:   omit.From(m.CreatedAt),
		Description: omitnull.FromNull(m.Description),
	}
	if m.ID > 0 {
		ss.ID = omit.From(m.ID)
	}
	return ss
}

type TagSetter struct {
	ID          omit.Val[int64]      `db:"id"`
	GUID        omit.Val[string]     `db:"guid"`
	Code        omit.Val[string]     `db:"code"`
	Name        omit.Val[string]     `db:"name"`
	CreatedAt   omit.Val[time.Time]  `db:"created_at"`
	Description omitnull.Val[string] `db:"description"`
}

func (t TAGS) Mapper() sqx.Mapper[TAGS, Tag, TagSetter] {
	return TagMapper{t: t}
}

func (t TAGS) Executor() sqx.Executor[Tag, TagSetter] {
	return sqx.NewExecutor(t.Mapper())
}

func (t TAGS) ColumnSetter(ctx context.Context, c *sq.Column, s TagSetter) {
	_ = ctx
	s.ID.IfSet(func(v int64) {
		c.SetInt64(t.ID, v)
	})
	s.GUID.IfSet(func(v string) {
		c.SetString(t.GUID, v)
	})
	s.Name.IfSet(func(v string) {
		c.SetString(t.NAME, v)
	})
	s.Code.IfSet(func(v string) {
		c.SetString(t.CODE, v)
	})
	s.CreatedAt.IfSet(func(v time.Time) {
		c.SetTime(t.CREATED_AT, v)
	})
	s.Description.IfSet(func(v string) {
		c.SetString(t.DESCRIPTION, v)
	})
}

func (t TAGS) RowSetter(r *sq.Row) Tag {
	v := Tag{}
	v.ID = r.Int64Field(t.ID)
	v.GUID = r.StringField(t.GUID)
	v.Name = r.StringField(t.NAME)
	v.Code = r.StringField(t.CODE)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	desc := r.NullStringField(t.DESCRIPTION)
	v.Description = null.FromCond(desc.String, desc.Valid)
	return v
}

type TagMapper struct {
	t TAGS
}

func (m TagMapper) Table() TAGS {
	return m.t
}

func (m TagMapper) InsertT(ctx context.Context, vv ...TagSetter) func(*sq.Column) {
	return func(c *sq.Column) {
		for _, v := range vv {
			m.t.ColumnSetter(ctx, c, v)
		}
	}
}

func (m TagMapper) UpdateT(ctx context.Context, v TagSetter) func(*sq.Column) {
	return func(c *sq.Column) {
		m.t.ColumnSetter(ctx, c, v)
	}
}

func (m TagMapper) SelectT(ctx context.Context) func(*sq.Row) Tag {
	return m.t.RowSetter
}

func (t TAGS) Insert(ctx context.Context, db sq.DB, ss ...TagSetter) (sq.Result, error) {
	q := sq.InsertInto(t).
		ColumnValues(t.InsertQ(ctx, ss...))
	return sq.Exec(db, q)
}

func (t TAGS) Update(ctx context.Context, db sq.DB, where sq.Predicate, s TagSetter) (sq.Result, error) {
	q := sq.Update(t).
		SetFunc(t.UpdateQ(ctx, s)).
		Where(where)
	return sq.Exec(db, q)
}

func (t TAGS) Delete(ctx context.Context, db sq.DB, where sq.Predicate) (sq.Result, error) {
	q := sq.DeleteFrom(t).
		Where(where)
	return sq.Exec(db, q)
}

func (t TAGS) One(ctx context.Context, db sq.DB, where sq.Predicate) (Tag, error) {
	q := sq.From(t).Where(where)
	row, err := sq.FetchOne(db, q, t.SelectQ(ctx))
	return row, err
}

func (t TAGS) All(ctx context.Context, db sq.DB, where sq.Predicate) ([]Tag, error) {
	q := sq.From(t).Where(where)
	rows, err := sq.FetchAll(db, q, t.SelectQ(ctx))
	return rows, err
}

func (t TAGS) InsertQ(ctx context.Context, ss ...TagSetter) func(*sq.Column) {
	q := func(c *sq.Column) {
		for _, s := range ss {
			t.ColumnSetter(ctx, c, s)
		}
	}
	return q
}

func (t TAGS) UpdateQ(ctx context.Context, s TagSetter) func(*sq.Column) {
	q := func(c *sq.Column) {
		t.ColumnSetter(ctx, c, s)
	}
	return q
}

func (t TAGS) SelectQ(ctx context.Context) func(*sq.Row) Tag {
	return func(r *sq.Row) Tag {
		return t.RowSetter(r)
	}
}

func (t TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().Eq(id)
}
