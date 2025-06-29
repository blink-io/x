package orm

import (
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

	ua := [][16]byte{
		uuid.New(),
		uuid.New(),
		uuid.New(),
	}

	_ = ua
	_ = j1

	d1 := ArraySetter{
		//ID:         omit.From(int64(gofakeit.Int32())),
		StrArrays:  omit.From([]string{"A", "B", "C"}),
		Int4Arrays: omit.From([]int32{1, 2, 3, 4, 5}),
		BoolArrays: omit.From([]bool{true, false, true}),
		CreatedAt:  omit.From(time.Now()),
		VJson:      omitnull.From(j1),
		VUUID:      omitnull.From([16]byte(uuid.New())),
		JsonArrays: omitnull.From([]map[string]any{
			j1,
		}),
		//UuidArrays: omitnull.From(ua),
	}

	var exec = Executors.Array
	_, err := exec.Insert(ctx, db, []ArraySetter{d1}...)
	require.NoError(t, err)
}
