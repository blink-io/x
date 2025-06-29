package orm

import (
	"context"
	"encoding/json"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	rawsq "github.com/bokwoon95/sq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

type RawARRAYS struct {
	rawsq.TableStruct
	ID           rawsq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	STR_ARRAYS   rawsq.ArrayField  `ddl:"type=varchar[] notnull"`
	INT4_ARRAYS  rawsq.ArrayField  `ddl:"type=int[] notnull"`
	BOOL_ARRAYS  rawsq.ArrayField  `ddl:"type=boolean[] notnull"`
	CREATED_AT   rawsq.TimeField   `ddl:"type=timestamptz notnull"`
	V_JSONB      rawsq.JSONField   `ddl:"type=jsonb"`
	V_JSON       rawsq.JSONField   `ddl:"type=json"`
	V_UUID       rawsq.UUIDField   `ddl:"type=uuid"`
	JSONB_ARRAYS rawsq.ArrayField  `ddl:"type=jsonb[]"`
	JSON_ARRAYS  rawsq.ArrayField  `ddl:"type=json[]"`
	UUID_ARRAYS  rawsq.ArrayField  `ddl:"type=uuid[]"`
	INT_AAA      rawsq.ArrayField  `ddl:"type=int[]"`
}

func init() {
	d := rawsq.DialectPostgres
	rawsq.DefaultDialect.Store(&d)
}

func (t RawARRAYS) ColumnSetter(ctx context.Context, c *rawsq.Column, ss ...ArraySetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.StrArrays.IfSet(func(v []string) {
			c.SetArray(t.STR_ARRAYS, v)
		})
		s.Int4Arrays.IfSet(func(v []int32) {
			c.SetArray(t.INT4_ARRAYS, v)
		})
		s.BoolArrays.IfSet(func(v []bool) {
			c.SetArray(t.BOOL_ARRAYS, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.VJsonb.IfSet(func(v map[string]any) {
			c.SetJSON(t.V_JSONB, v)
		})
		s.VJson.IfSet(func(v map[string]any) {
			c.SetJSON(t.V_JSON, v)
		})
		s.VUUID.IfSet(func(v [16]byte) {
			c.SetUUID(t.V_UUID, v)
		})
		s.JsonbArrays.IfSet(func(v []map[string]any) {
			var jonsStrArrays []string
			for _, iv := range v {
				var b strings.Builder
				if err := json.NewEncoder(&b).Encode(iv); err != nil {
					panic(err)
				}
				jonsStrArrays = append(jonsStrArrays, b.String())
			}
			c.SetArray(t.JSONB_ARRAYS, jonsStrArrays)
		})
		s.JsonArrays.IfSet(func(v []map[string]any) {
			var jonsStrArrays []string
			for _, iv := range v {
				var b strings.Builder
				if err := json.NewEncoder(&b).Encode(iv); err != nil {
					panic(err)
				}
				jonsStrArrays = append(jonsStrArrays, b.String())
			}
			c.SetArray(t.JSON_ARRAYS, jonsStrArrays)
		})
		s.UuidArrays.IfSet(func(v [][16]byte) {
			var jonsStrArrays []string
			for _, iv := range v {
				uv, err := uuid.FromBytes(iv[:])
				if err != nil {
					panic(err)
				}
				jonsStrArrays = append(jonsStrArrays, uv.String())
			}
			c.SetArray(t.UUID_ARRAYS, jonsStrArrays)
		})
		s.IntAaa.IfSet(func(v []int32) {
			c.SetArray(t.INT_AAA, v)
		})
	}
}

func (t RawARRAYS) ColumnMapper(ss ...ArraySetter) func(*rawsq.Column) {
	return func(c *rawsq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func TestPg_Arrays_RawInsert_1(t *testing.T) {
	db := GetPgDB()
	var rawArrays = rawsq.New[RawARRAYS]("")

	j1 := map[string]any{
		"foo": "bar",
		"bar": "baz",
	}

	_ = j1
	//ua := [][16]byte{
	//	uuid.New(),
	//	uuid.New(),
	//	uuid.New(),
	//}

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
		UuidArrays: omitnull.From([][16]byte{
			uuid.New(),
		}),
	}
	var q = rawsq.InsertInto(Tables.Arrays).
		ColumnValues(rawArrays.ColumnMapper(d1))
	_, err := rawsq.Exec(rawsq.VerboseLog(db), q)
	require.NoError(t, err)
}
