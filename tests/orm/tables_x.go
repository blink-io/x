package orm

import "github.com/bokwoon95/sq"

func (s TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{s.ID}
}

func (s TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return s.PrimaryKeys().Eq(id)
}

var Tvals = sq.New[TVALS]("tt1")

func (s TVALS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{s.IID, s.SID}
}

func (s TVALS) PrimaryKeyValues(iid int64, sid string) sq.Predicate {
	return s.PrimaryKeys().In(sq.RowValues{{iid, sid}})
}
