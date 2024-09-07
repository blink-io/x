package orm

import (
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Mysql_Mkey_Insert_1(t *testing.T) {
	db := getMysqlDBForSQ()
	tbl := sq.New[MKEYS]("")

	prefix := "from-sq:"

	ss := sq.InsertInto(tbl).
		Columns(tbl.ID1, tbl.ID2, tbl.GUID, tbl.NAME, tbl.CREATED_AT).
		Values(11, 12, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(13, 14, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(23, 24, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now())
	ss.RowAlias = ""

	_, err := sq.Exec(sq.VerboseLog(db), ss)
	require.NoError(t, err)
}

func TestSq_Mysql_Mkey_FetchOne_ByID(t *testing.T) {
	db := getMysqlDBForSQ()
	idb := sq.VerboseLog(db)
	tbl := Tables.Mkeys

	query := sq.Select(tbl.ID1, tbl.ID2, tbl.NAME).
		From(tbl).Where(
		tbl.PrimaryKeys().In(sq.RowValues{{11, 12}})).
		Limit(100)

	type result struct {
		ID1  int
		ID2  int
		Name string
	}
	records, err := sq.FetchAll(idb, query, func(r *sq.Row) result {
		return result{
			ID1:  r.Int("id1"),
			ID2:  r.Int("id2"),
			Name: r.String("name"),
			//GUID:      r.StringField(tbl.GUID),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Mysql_Mkey_FetchAll_1(t *testing.T) {
	db := getMysqlDBForSQ()
	ldb := sq.VerboseLog(db)
	tbl := Tables.Mkeys

	where := sq.RowValue{tbl.ID1}.Eq(11)
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
