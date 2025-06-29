package orm

import (
	"testing"

	"github.com/blink-io/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Sqlite_Tag_Insert_1(t *testing.T) {
	db := GetSqliteDB()
	sb := sq.SQLite
	tbl := Tables.Tags

	ss := []TagSetter{
		randomTagSetter(nil),
		randomTagSetter(Ptr(gofakeit.City())),
		randomTagSetter(nil),
	}

	rt, err := sq.Exec(
		sq.Log(db),
		sb.InsertInto(tbl).
			ColumnValues(tbl.ColumnMapper(ss...)),
	)

	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Sqlite_Tag_Insert_2(t *testing.T) {
	db := GetSqliteDB()

	ss := []TagSetter{
		randomTagSetter(nil),
		randomTagSetter(Ptr(gofakeit.City())),
	}

	rt, err := Executors.Tag.Insert(ctx, sq.Log(db), ss...)
	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Sqlite_Tag_Insert_Select_1(t *testing.T) {
	db := GetSqliteDB()
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
	db := GetSqliteDB()
	tbl := Tables.Tags

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, tbl.RowMapper)

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite_Tag_Mapper_FetchAll_2(t *testing.T) {
	db := GetSqliteDB()
	tbl := Tables.Tags

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, tbl.RowMapperFunc())

	require.NoError(t, err)
	require.NotNil(t, records)
}
