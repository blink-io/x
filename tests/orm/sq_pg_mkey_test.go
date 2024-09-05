package orm

import (
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

type MKEYS struct {
	sq.TableStruct `ddl:"primarykey=id1,id2"`
	ID1            sq.NumberField `ddl:"type=int notnull"`
	ID2            sq.NumberField `ddl:"type=int notnull"`
	NAME           sq.StringField `ddl:"type=varchar(60) notnull"`
	CREATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
	GUID           sq.StringField `ddl:"type=varchar(60) notnull unique"`
}

func (s MKEYS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{s.ID1, s.ID2}
}

func (s MKEYS) PrimaryKeyValues(id1, id2 int64) sq.Predicate {
	return s.PrimaryKeys().In(sq.RowValues{{id1, id2}})
}

type Mkey struct {
	ID1       int       `db:"id1"`
	ID2       int       `db:"id2"`
	GUID      string    `db:"guid"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

var MkeyTable = sq.New[MKEYS]("aa")

func TestSq_Pg_Mkey_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := MkeyTable

	prefix := "from-sq:"

	_, err := sq.Exec(sq.VerboseLog(db), sq.
		InsertInto(tbl).
		Columns(tbl.ID1, tbl.ID2, tbl.GUID, tbl.NAME, tbl.CREATED_AT).
		Values(11, 12, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(13, 14, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()).
		Values(23, 24, gofakeit.UUID(), prefix+gofakeit.Username(), time.Now()),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Mkey_FetchOne_ByID(t *testing.T) {
	db := getPgDBForSQ()
	idb := sq.VerboseLog(db)
	tbl := MkeyTable

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
			ID1:  r.Int("id1"),
			ID2:  r.Int("id2"),
			Name: r.String("name"),
			GUID: r.String("guid"),
			//CreatedAt: r.TimeField(tbl.CREATED_AT),
		}
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Mkey_FetchAll_1(t *testing.T) {
	db := getPgDBForSQ()
	ldb := sq.VerboseLog(db)
	tbl := MkeyTable

	where := MkeyTable.PrimaryKeys().In(11)
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
