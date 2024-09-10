package orm

import (
	"github.com/bokwoon95/sq"
)

type tables struct {
	Tags   TAGS
	Users  USERS
	Tvals  TVALS
	Mkeys  MKEYS
	Arrays ARRAYS
}

var Tables = tables{
	Tags:   sq.New[TAGS](""),
	Users:  sq.New[USERS](""),
	Tvals:  sq.New[TVALS](""),
	Mkeys:  sq.New[MKEYS](""),
	Arrays: sq.New[ARRAYS](""),
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
