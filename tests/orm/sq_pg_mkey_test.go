package orm

import (
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Pg_Mkey_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Mkeys

	prefix := "from-sq:"

	_, err := sq.Exec(sq.VerboseLog(db), sq.
		InsertInto(tbl).
		Columns(tbl.GUID, tbl.NAME, tbl.CREATED_AT).
		Values(gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Mkey_FetchOne_ByID(t *testing.T) {
	db := getPgDBForSQ()
	idb := sq.VerboseLog(db)
	tbl := Tables.Mkeys

	idWhere := tbl.PrimaryKeyValues(11, 12)
	query := sq.Select(tbl.ID1, tbl.ID2, tbl.NAME, tbl.GUID).
		From(tbl).Where(idWhere).
		Limit(100)

	type result struct {
		ID1  int
		ID2  int
		GUID string
		Name string
	}
	records, err := sq.FetchAll(idb, query, func(r *sq.Row) result {
		return result{
			ID1:  r.Int(tbl.ID1.GetName()),
			ID2:  r.Int(tbl.ID2.GetName()),
			Name: r.String(tbl.NAME.GetName()),
			GUID: r.String(tbl.GUID.GetName()),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Mkey_FetchAll_1(t *testing.T) {
	db := getPgDBForSQ()
	ldb := sq.VerboseLog(db)
	tbl := Tables.Mkeys

	where := tbl.PrimaryKeys().In(11)
	query := sq.From(tbl).Where(where).
		Limit(100)

	records, err := sq.FetchAll(ldb, query, func(r *sq.Row) Mkey {
		return Mkey{
			ID1:       r.IntField(tbl.ID1),
			ID2:       r.IntField(tbl.ID2),
			Name:      r.StringField(tbl.NAME),
			GUID:      r.StringField(tbl.GUID),
			CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}
