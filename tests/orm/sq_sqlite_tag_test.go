package orm

import (
	"testing"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Sqlite_Tag_Insert_1(t *testing.T) {
	db := getSqliteDBForSQ()
	sb := sq.SQLite
	tbl := Tables.Tags

	records := []Tag{
		randomTag(nil),
		randomTag(Ptr(gofakeit.City())),
		randomTag(nil),
	}

	rt, err := sq.Exec(sq.Log(db), sb.
		InsertInto(tbl).
		ColumnValues(func(col *sq.Column) {
			for _, r := range records {
				tagInsertColumnMapper(col, r)
			}
		}))

	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Sqlite_Tag_Insert_2(t *testing.T) {
	db := getSqliteDBForSQ()

	err := randomTag(nil).Insert(db)
	require.NoError(t, err)

	err = randomTag(Ptr(gofakeit.School())).Insert(db)
	require.NoError(t, err)
}

func TestSq_Sqlite_Tag_Insert_Select_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := Tables.TagsBak
	fTbl := Tables.Tags

	_, err := sq.Exec(db, sq.Postgres.
		InsertInto(tbl).
		Columns(tbl.GUID, tbl.NAME, tbl.CODE, tbl.DESCRIPTION, tbl.CREATED_AT).
		Select(sq.
			Select(fTbl.GUID, fTbl.NAME, fTbl.CODE, fTbl.DESCRIPTION, fTbl.CREATED_AT).
			From(fTbl).
			Where(fTbl.DESCRIPTION.IsNotNull()),
		).
		SetDialect(sq.DialectPostgres),
	)
	require.NoError(t, err)
}

func TestSq_Sqlite_Tag_Mapper_FetchAll_1(t *testing.T) {
	db := getSqliteDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, mm.QueryT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite_Tag_Mapper_FetchAll_2(t *testing.T) {
	db := getSqliteDBForSQ()
	mm := Mappers.TAGS
	tbl := mm.Table()

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, mm.QueryT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}
