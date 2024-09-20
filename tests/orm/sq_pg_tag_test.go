package orm

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/ptr"
	"github.com/blink-io/x/sql/misc"
	"github.com/blink-io/x/types/tuplen"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Pg_Tag_Insert_1(t *testing.T) {
	db := getPgDBForSQ()

	tbl := TagTable

	records := []Tag{
		randomTag(nil),
		randomTag(Ptr(gofakeit.City())),
		randomTag(nil),
		randomTag(nil),
		randomTag(nil),
		randomTag(nil),
		randomTag(nil),
		randomTag(nil),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			tagInsertColumnMapper(col, r)
		}
	}))

	require.NoError(t, err)
}

func TestSq_Pg_Tag_Insert_2(t *testing.T) {
	db := getPgDBForSQ()

	err := randomTag(nil).Insert(db)
	require.NoError(t, err)

	err = randomTag(Ptr(gofakeit.School())).Insert(db)
	require.NoError(t, err)

}

func TestSq_Pg_Tag_Mapper_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := NewTagMapper()
	tbl := mm.Table()

	d1 := randomTag(nil)
	d2 := randomTag(ptr.Of("Hello, Hi, 你好"))

	_, err := sq.Exec(sq.Log(db), sq.
		InsertInto(tbl).
		ColumnValues(mm.InsertMapper(d1, d2)),
	)
	require.NoError(t, err)
}

func TestSq_Pg_Tag_Mapper_FetchAll_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := NewTagMapper()
	tbl := mm.Table()

	query := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, query, mm.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Tag_Fetch_Custom_1(t *testing.T) {
	db := getPgDBForSQ()
	mm := NewTagMapper()
	tbl := mm.Table()

	query := sq.
		Queryf("select id code, name from tags limit 5")
	records, err := sq.FetchAll(db, query, func(r *sq.Row) tuplen.Tuple3[int, string, string] {
		return tuplen.Of3(
			r.Int(tbl.ID.GetName()),
			r.String(tbl.CODE.GetName()),
			r.String(tbl.NAME.GetName()),
		)
	})

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Tag_Mapper_FetchAll_Paging(t *testing.T) {
	db := getPgDBForSQ()
	mm := NewTagMapper()
	tbl := mm.Table()
	perPage := 3
	qm := mm.QueryMapper()
	vdb := sq.VerboseLog(db)

	bq := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(perPage).
		OrderBy(tbl.ID.Asc())

	t.Run("pagination", func(t *testing.T) {
		q1 := bq.Offset(misc.ToOffset(1, perPage))
		rs1, err1 := sq.FetchAll(vdb, q1, qm)
		require.NoError(t, err1)
		require.NotNil(t, rs1)

		q2 := bq.Offset(misc.ToOffset(2, perPage))
		rs2, err2 := sq.FetchAll(vdb, q2, qm)
		require.NoError(t, err2)
		require.NotNil(t, rs2)

		q3 := bq.Offset(misc.ToOffset(3, perPage))
		rs3, err3 := sq.FetchAll(vdb, q3, qm)
		require.NoError(t, err3)
		require.NotNil(t, rs3)
	})

	fmt.Println("done")
}

func TestSq_Pg_Tag_Mapper_FetchOne_ByPK(t *testing.T) {
	db := getPgDBForSQ()
	mm := NewTagMapper()
	tbl := mm.Table()

	idWhere := tbl.PrimaryKeyValues(23)
	query := sq.
		From(tbl).
		Where(idWhere)
	records, err := sq.FetchOne(db, query, mm.QueryMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func tagInsertColumnMapper(c *sq.Column, r Tag) {
	tbl := Tables.Tags

	c.SetString(tbl.GUID, r.GUID)
	c.SetString(tbl.NAME, r.Name)
	c.SetString(tbl.CODE, r.Code)
	c.Set(tbl.DESCRIPTION, r.Description)
	c.SetTime(tbl.CREATED_AT, r.CreatedAt)
}
