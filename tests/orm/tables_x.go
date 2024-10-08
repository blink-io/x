package orm

import (
	"context"
	"time"

	"github.com/bokwoon95/sq"
)

type tables struct {
	Tags    TAGS
	TagsBak TAGS_BAK
	Users   USERS
	Tvals   TVALS
	Mkeys   MKEYS
	Arrays  ARRAYS
}

var Tables = tables{
	Tags:    sq.New[TAGS](""),
	TagsBak: sq.New[TAGS_BAK](""),
	Users:   sq.New[USERS](""),
	Tvals:   sq.New[TVALS](""),
	Mkeys:   sq.New[MKEYS](""),
	Arrays:  sq.New[ARRAYS](""),
}

func (t TAGS) Insert(db sq.DB, ms ...Tag) (sq.Result, error) {
	q := sq.InsertInto(t).ColumnValues(func(col *sq.Column) {
		for _, m := range ms {
			col.SetString(t.GUID, m.GUID)
			col.SetString(t.NAME, m.Name)
			col.SetString(t.CODE, m.Code)
			col.SetTime(t.CREATED_AT, m.CreatedAt)
			col.SetString(t.DESCRIPTION, m.Description.GetOr(""))
		}
	})
	return sq.Exec(db, q)
}

func (t TAGS) Insert2(db sq.DB, ms ...TagSetter) (sq.Result, error) {
	q := sq.InsertInto(t).ColumnValues(func(c *sq.Column) {
		for _, m := range ms {
			m.ID.IfSet(func(v int) {
				c.SetInt(t.ID, v)
			})
			m.GUID.IfSet(func(v string) {
				c.SetString(t.GUID, v)
			})
			m.Name.IfSet(func(v string) {
				c.SetString(t.NAME, v)
			})
			m.Code.IfSet(func(v string) {
				c.SetString(t.CODE, v)
			})
			m.CreatedAt.IfSet(func(v time.Time) {
				c.SetTime(t.CREATED_AT, v)
			})
			c.SetString(t.DESCRIPTION, m.Description.GetOr(""))
		}
	})
	return sq.Exec(db, q)
}

func (t TAGS) Update(db sq.DB, m TagSetter) (sq.Result, error) {
	q := sq.Update(t).SetFunc(func(c *sq.Column) {
		if !m.ID.IsUnset() {
			v, _ := m.ID.Get()
			c.SetInt(t.ID, v)
		}
		if !m.GUID.IsUnset() {
			v, _ := m.GUID.Get()
			c.SetString(t.GUID, v)
		}
		if !m.Name.IsUnset() {
			v, _ := m.Name.Get()
			c.SetString(t.NAME, v)
		}
		if !m.Code.IsUnset() {
			v, _ := m.Code.Get()
			c.SetString(t.CODE, v)
		}
		if !m.Description.IsUnset() {
			v, _ := m.Description.Get()
			c.SetString(t.DESCRIPTION, v)
		}
	})
	return sq.Exec(db, q)
}

func (t TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().Eq(id)
}

func (t TVALS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.IID, t.SID}
}

func (s TVALS) PrimaryKeyValues(iid int64, sid string) sq.Predicate {
	return s.PrimaryKeys().Eq(sq.RowValues{{iid, sid}})
}

var alwaysTrueExpr = sq.Expr("1 = {}", 1)

const (
	defaultTenantID = 1
)

func (tbl USERS) Policy(ctx context.Context, dialect string) (sq.Predicate, error) {
	tenantID, ok := ctx.Value(tbl.TENANT_ID.GetName()).(int)
	if !ok {
		return sq.And(alwaysTrueExpr, tbl.TENANT_ID.EqInt(defaultTenantID)), nil
	}
	return tbl.TENANT_ID.EqInt(tenantID), nil
}

func (tbl USERS) InsertMapper(rs ...User) func(c *sq.Column) {
	return func(c *sq.Column) {
		for _, r := range rs {
			if r.ID > 0 {
				c.SetInt(tbl.ID, r.ID)
			}
			c.SetString(tbl.GUID, r.GUID)
			c.SetString(tbl.USERNAME, r.Username)
			c.SetString(tbl.FIRST_NAME, r.FirstName)
			c.SetString(tbl.LAST_NAME, r.LastName)
			c.SetFloat64(tbl.SCORE, r.Score)
			c.SetInt(tbl.LEVEL, r.Level)
			c.SetTime(tbl.CREATED_AT, r.CreatedAt)
			c.SetTime(tbl.UPDATED_AT, r.UpdatedAt)
			c.SetInt(tbl.TENANT_ID, r.TenantID)
		}
	}
}

func (tbl USERS) QueryMapper() func(*sq.Row) *User {
	return func(r *sq.Row) *User {
		u := &User{
			ID:        r.IntField(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			Username:  r.StringField(tbl.USERNAME),
			FirstName: r.StringField(tbl.FIRST_NAME),
			LastName:  r.StringField(tbl.LAST_NAME),
			Score:     r.Float64Field(tbl.SCORE),
			Level:     r.IntField(tbl.LEVEL),
			TenantID:  r.IntField(tbl.TENANT_ID),
		}

		dd := sq.DefaultDialect.Load()
		dz := sq.DialectSQLite
		if dz == *dd {
			crstr := r.String("created_at")
			upstr := r.String("updated_at")
			if ct, err := time.Parse(time.RFC3339Nano, crstr); err == nil {
				u.CreatedAt = ct
			}
			if ct, err := time.Parse(time.RFC3339Nano, upstr); err == nil {
				u.UpdatedAt = ct
			}
		} else {
			u.CreatedAt = r.TimeField(tbl.CREATED_AT)
			u.UpdatedAt = r.TimeField(tbl.UPDATED_AT)
		}

		return u
	}
}
