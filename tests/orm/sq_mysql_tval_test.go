package orm

import (
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Mysql_TVAL_Insert_1(t *testing.T) {
	db := getMysqlDBForSQ()
	tbl := sq.New[TVALS]("")

	prefix := "from-sq:"

	ss := sq.InsertInto(tbl).
		Columns(tbl.IID, tbl.SID, tbl.NAME, tbl.CREATED_AT).
		Values(gofakeit.Uint8(), gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(gofakeit.Uint8(), gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(gofakeit.Uint8(), gofakeit.UUID(), prefix+gofakeit.Username(), time.Now())
	ss.RowAlias = ""

	rt, err := sq.Exec(sq.VerboseLog(db), ss)
	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Mysql_TVAL_FetchOne_ByID(t *testing.T) {
	db := getMysqlDBForSQ()
	idb := sq.VerboseLog(db)
	tbl := Tables.Tvals

	idWhere := tbl.PrimaryKeyValues(36, "9ea443a7-ac20-44ad-881a-28578e92250d")
	query := sq.Select(tbl.IID, tbl.SID, tbl.NAME).
		From(tbl).
		Where(idWhere, idWhere, tbl.NAME.LikeString("from-sq%"))

	type result struct {
		ID1  int
		ID2  string
		Name string
	}
	records, err := sq.FetchOne(idb, query, func(r *sq.Row) result {
		return result{
			ID1:  r.Int("iid"),
			ID2:  r.String("sid"),
			Name: r.String("name"),
			//UserGUID:      r.StringField(tbl.UserGUID),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Mysql_TVAL_FetchAll_1(t *testing.T) {
	db := getMysqlDBForSQ()
	ldb := sq.VerboseLog(db)
	tbl := Tables.Tvals

	//idWhere := tbl.PrimaryKeyValues(36, "9ea443a7-ac20-44ad-881a-28578e92250d")
	iidGtZero := tbl.IID.GtInt64(0)
	query := sq.From(tbl).Where(
		iidGtZero,
		tbl.NAME.LikeString("%Bruen%"),
	).
		OrderBy(tbl.SID.Desc()).
		Limit(100)

	records, err := sq.FetchAll(ldb, query, func(r *sq.Row) map[string]any {
		return map[string]any{
			"IID":       r.IntField(tbl.IID),
			"SID":       r.StringField(tbl.SID),
			"Name":      r.StringField(tbl.NAME),
			"CreatedAt": r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}
