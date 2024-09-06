package orm

import "github.com/bokwoon95/sq"

type tables struct {
	Tags  TAGS
	Users USERS
	Tvals TVALS
	Mkeys MKEYS
}

var Tables = tables{
	Tags:  sq.New[TAGS](""),
	Users: sq.New[USERS](""),
	Tvals: sq.New[TVALS](""),
	Mkeys: sq.New[MKEYS](""),
}

func (s TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{s.ID}
}

func (s TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return s.PrimaryKeys().Eq(id)
}

func (s TVALS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{s.IID, s.SID}
}

func (s TVALS) PrimaryKeyValues(iid int64, sid string) sq.Predicate {
	return s.PrimaryKeys().Eq(sq.RowValues{{iid, sid}})
}
