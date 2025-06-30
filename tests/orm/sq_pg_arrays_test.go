package orm

import (
	"math"
	"testing"
	"time"

	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
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
	results, err := exec.All(ctx, db, where)
	require.NoError(t, err)
	require.True(t, len(results) > 0)
}
