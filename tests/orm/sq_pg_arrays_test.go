package orm

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/blink-io/sq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPg_Arrays_Insert_1(t *testing.T) {
	db := GetPgDB()

	j1 := map[string]any{
		"foo": "bar",
		"bar": "baz",
	}

	d1 := ArraySetter{
		//ID:         omit.From(int64(gofakeit.Int32())),
		StrArrays:  omit.From([]string{"A", "B", "C"}),
		Int4Arrays: omit.From([]int32{1, 2, 3, 4, 5, math.MaxInt32}),
		Int2Arrays: omitnull.From([]int16{math.MaxInt16}),
		BoolArrays: omit.From([]bool{true, false, true}),
		CreatedAt:  omit.From(time.Now()),
		VJson:      omitnull.From(j1),
		VUUID:      omitnull.From([16]byte(uuid.New())),
		JsonArrays: omitnull.From([]string{
			`{"foo": "bar1111"}`,
			`{"foo": "bar2222"}`,
		}),
		UuidArrays: omitnull.From([]string{
			uuid.New().String(),
			uuid.New().String(),
			uuid.New().String(),
		}),
		TsArrays: omitnull.From([]string{
			time.Now().Format(time.RFC3339),
			time.Now().Format(time.RFC3339),
		}),
	}

	var exec = Executors.Array
	_, err := exec.Insert(ctx, db, []ArraySetter{d1}...)
	require.NoError(t, err)
}

func TestSelectAll_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Array
	where := exec.Table().ID.GtInt(0)
	results, err := exec.All(ctx, sq.VerboseLog(db), where)
	require.NoError(t, err)
	require.True(t, len(results) > 0)
}

func TestSelectCustom_JSON_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Array
	tbl := exec.Table()

	where := tbl.V_JSON.IsNotNull()
	q := sq.Select(tbl.ID, tbl.V_JSON).From(tbl).Where(where)
	rr, err := sq.FetchAll[map[string]any](db, q, func(ctx context.Context, row *sq.Row) map[string]any {
		var mm map[string]any
		row.JSON(&mm, tbl.V_JSON.GetName())
		return mm
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}

func TestSelectCustom_UUID_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Array
	tbl := exec.Table()

	where := tbl.V_UUID.IsNotNull()
	q := sq.Select(tbl.ID, tbl.V_UUID).From(tbl).Where(where)
	rr, err := sq.FetchAll[[16]byte](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) [16]byte {
		var mm [16]byte
		row.UUID(&mm, tbl.V_UUID.GetName())
		return mm
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)

	for _, b := range rr {
		u := uuid.UUID(b)
		require.NoError(t, err)
		fmt.Println(u.String())
	}
}

func TestSelectCustom_Int4Array_1(t *testing.T) {
	db := GetPgDB()
	var exec = Executors.Array
	tbl := exec.Table()

	where := tbl.INT4_ARRAYS.IsNotNull()
	q := sq.Select(tbl.ID, tbl.REMARK,
		tbl.STR_ARRAYS, tbl.BOOL_ARRAYS,
		tbl.JSON_ARRAYS, tbl.INT4_ARRAYS).From(tbl).Where(where)
	rr, err := sq.FetchAll[Array](sq.VerboseLog(db), q, func(ctx context.Context, row *sq.Row) Array {
		var r Array
		int32Arrays := new([]int32)
		strArrays := new([]string)
		jsonArrays := new([]string)
		boolArrays := new([]bool)
		row.Array(int32Arrays, tbl.INT4_ARRAYS.GetName())
		row.Array(strArrays, tbl.STR_ARRAYS.GetName())
		row.Array(boolArrays, tbl.BOOL_ARRAYS.GetName())
		row.Array(jsonArrays, tbl.JSON_ARRAYS.GetName())

		r.Int4Arrays = *int32Arrays
		r.BoolArrays = *boolArrays
		r.StrArrays = *strArrays
		r.JsonArrays = null.FromPtr(jsonArrays)

		var id = new(int64)
		row.Scan(id, tbl.ID.GetName())

		var remark = new(string)
		row.Scan(remark, tbl.REMARK.GetName())
		fmt.Printf("remark:%s\n", *remark)

		var nullRemark sq.NullString
		row.Scan(&nullRemark, tbl.REMARK.GetName())
		r.ID = *id
		return r
	})
	require.NoError(t, err)
	require.True(t, len(rr) > 0)
}
