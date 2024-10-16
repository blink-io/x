package orm

import (
	"testing"

	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Sqlite_Tag_Insert_1(t *testing.T) {
	db := GetSqliteDB()
	sb := sq.SQLite
	tbl := Tables.Tags

	ss := []TagSetter{
		randomTag(nil).Setter(),
		randomTag(Ptr(gofakeit.City())).Setter(),
		randomTag(nil).Setter(),
	}

	rt, err := sq.Exec(
		sq.Log(db),
		sb.InsertInto(tbl).
			ColumnValues(tbl.InsertQ(ctx, ss...)),
	)

	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Sqlite_Tag_Insert_2(t *testing.T) {
	db := GetSqliteDB()
	//sb := sq.SQLite
	tbl := Tables.Tags

	s1 := randomTag(nil).Setter()
	s2 := randomTag(Ptr(gofakeit.School())).Setter()

	rt, err := tbl.Insert(ctx, sq.Log(db), s1, s2)
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
	mm := Mappers.TAGS
	tbl := mm.Table()

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, mm.SelectT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite_Tag_Mapper_FetchAll_2(t *testing.T) {
	db := GetSqliteDB()
	mm := Mappers.TAGS
	tbl := mm.Table()

	q := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, q, mm.SelectT(ctx))

	require.NoError(t, err)
	require.NotNil(t, records)
}
