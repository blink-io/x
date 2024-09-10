package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/sql/misc"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestSq_Pg_Array_Insert_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Arrays

	records := []Array{
		randomArray(),
		randomArray(),
		randomArray(),
		randomArray(),
		randomArray(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			arrayInsertMapper(col, r)
		}
	}))

	require.NoError(t, err)
}

func TestSq_Pg_Array_FetchAll_1(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Arrays
	qm := arrayQueryMapper

	query := sq.
		From(tbl).
		Where(tbl.ID.GtInt(0)).
		Limit(100)
	records, err := sq.FetchAll(db, query, qm)

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Pg_Array_FetchAll_Paging(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Arrays
	perPage := 3
	qm := arrayQueryMapper
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

func TestSq_Pg_Array_FetchAll_2(t *testing.T) {
	db := getPgDBForSQ()
	tbl := Tables.Arrays
	qm := arrayQueryMapper

	query := sq.
		From(tbl)
	records, err := sq.FetchAll(db, query, qm)

	require.NoError(t, err)
	require.NotNil(t, records)
}

func randomArray() Array {
	return Array{
		CreatedAt: time.Now().Local(),

		StrArrays: []string{
			gofakeit.UUID(),
			gofakeit.DomainName(),
			gofakeit.AppName(),
		},

		Int4Arrays: []int32{
			gofakeit.Int32(),
			gofakeit.Int32(),
			gofakeit.Int32(),
			gofakeit.Int32(),
		},

		BoolArrays: []bool{
			gofakeit.Bool(),
			gofakeit.Bool(),
		},

		VJsonb: map[string]any{
			"country": gofakeit.Country(),
			"state":   gofakeit.State(),
			"city":    gofakeit.City(),
		},

		VJson: map[string]any{
			"dog":   gofakeit.Dog(),
			"cat":   gofakeit.Cat(),
			"fruit": gofakeit.Fruit(),
		},
	}
}

func arrayQueryMapper(r *sq.Row) Array {
	tbl := Tables.Arrays
	var strArrPtr []string
	var int4ArrPtr []int32
	var boolArrPtr []bool
	r.ArrayField(&strArrPtr, tbl.STR_ARRAYS)
	r.ArrayField(&int4ArrPtr, tbl.INT4_ARRAYS)
	r.ArrayField(&boolArrPtr, tbl.BOOL_ARRAYS)

	var vjsonb map[string]any
	var vjson map[string]any
	r.JSONField(&vjsonb, tbl.V_JSONB)
	r.JSONField(&vjson, tbl.V_JSON)

	return Array{
		StrArrays:  strArrPtr,
		Int4Arrays: int4ArrPtr,
		BoolArrays: boolArrPtr,

		VJsonb: vjsonb,
		VJson:  vjson,

		CreatedAt: r.TimeField(tbl.CREATED_AT),
	}
}

func arrayInsertMapper(c *sq.Column, r Array) {
	tbl := Tables.Arrays

	c.SetArray(tbl.STR_ARRAYS, r.StrArrays)
	c.SetArray(tbl.INT4_ARRAYS, r.Int4Arrays)
	c.SetArray(tbl.BOOL_ARRAYS, r.BoolArrays)
	c.SetTime(tbl.CREATED_AT, r.CreatedAt)

	c.SetJSON(tbl.V_JSONB, r.VJsonb)
	c.SetJSON(tbl.V_JSON, r.VJson)
}
