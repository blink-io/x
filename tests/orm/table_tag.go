package orm

import (
	"context"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/bokwoon95/sq"
)

func (t TAGS) setterToColumn(ctx context.Context, s TagSetter, c *sq.Column) {
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
	q := sq.From(t).Where(where).Limit(1)
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
			t.setterToColumn(ctx, s, c)
		}
	}
	return q
}

func (t TAGS) UpdateQ(ctx context.Context, s TagSetter) func(*sq.Column) {
	q := func(c *sq.Column) {
		t.setterToColumn(ctx, s, c)
	}
	return q
}

func (t TAGS) SelectQ(ctx context.Context) func(*sq.Row) Tag {
	return func(r *sq.Row) Tag {
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
}

func (t TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().Eq(id)
}
