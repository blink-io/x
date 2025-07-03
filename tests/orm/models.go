package orm

import (
	"context"
	"database/sql"
	"time"

	"github.com/blink-io/sq"
)

type Array struct {
	ID           int64                    ` db:"-" json:"-"`
	StrArrays    []string                 ` db:"str_arrays" json:"str_arrays"`
	Int4Arrays   []int32                  ` db:"int4_arrays" json:"int4_arrays"`
	BoolArrays   []bool                   ` db:"bool_arrays" json:"bool_arrays"`
	CreatedAt    time.Time                ` db:"created_at" json:"created_at"`
	VJsonb       sql.Null[map[string]any] ` db:"v_jsonb" json:"v_jsonb"`
	VJson        sql.Null[map[string]any] ` db:"v_json" json:"v_json"`
	VUUID        sql.Null[[16]byte]       ` db:"v_uuid" json:"v_uuid"`
	JsonbArrays  sql.Null[[]string]       ` db:"jsonb_arrays" json:"jsonb_arrays"` // []map[string]any
	JsonArrays   sql.Null[[]string]       ` db:"json_arrays" json:"json_arrays"`   // []map[string]any
	UuidArrays   sql.Null[[]string]       ` db:"uuid_arrays" json:"uuid_arrays"`   // [][16]byte
	IntAaa       sql.Null[[]int32]        ` db:"int_aaa" json:"int_aaa"`
	TsArrays     sql.Null[[]string]       ` db:"ts_arrays" json:"ts_arrays"` // []time.Time
	Int2Arrays   sql.Null[[]int16]        ` db:"int2_arrays" json:"int2_arrays"`
	Remark       sql.Null[string]         ` db:"remark" json:"remark"`
	StatusArrays sql.Null[[]any]          ` db:"status_arrays" json:"status_arrays"`
}

type ArraySetter struct {
	ID           *int64                    `db:"-" json:"-"`
	StrArrays    *[]string                 `db:"str_arrays" json:"str_arrays"`
	Int4Arrays   *[]int32                  `db:"int4_arrays" json:"int4_arrays"`
	BoolArrays   *[]bool                   `db:"bool_arrays" json:"bool_arrays"`
	CreatedAt    *time.Time                `db:"created_at" json:"created_at"`
	VJsonb       *sql.Null[map[string]any] `db:"v_jsonb" json:"v_jsonb"`
	VJson        *sql.Null[map[string]any] `db:"v_json" json:"v_json"`
	VUUID        *sql.Null[[16]byte]       `db:"v_uuid" json:"v_uuid"`
	JsonbArrays  *sql.Null[[]string]       `db:"jsonb_arrays" json:"jsonb_arrays"` // []map[string]any
	JsonArrays   *sql.Null[[]string]       `db:"json_arrays" json:"json_arrays"`   // []map[string]any
	UuidArrays   *sql.Null[[]string]       `db:"uuid_arrays" json:"uuid_arrays"`   // [][16]byte
	IntAaa       *sql.Null[[]int32]        `db:"int_aaa" json:"int_aaa"`
	TsArrays     *sql.Null[[]string]       `db:"ts_arrays" json:"ts_arrays"` // []time.Time
	Int2Arrays   *sql.Null[[]int16]        `db:"int2_arrays" json:"int2_arrays"`
	Remark       *sql.Null[string]         `db:"remark" json:"remark"`
	StatusArrays *sql.Null[[]any]          `db:"status_arrays" json:"status_arrays"`
}

func (t ARRAYS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...ArraySetter) {
	for idx, s := range ss {
		_ = idx
		if s.ID != nil {
			c.SetInt64(t.ID, *s.ID)
		}
		if s.StrArrays != nil {
			c.SetArray(t.STR_ARRAYS, *s.StrArrays)
		}
		if s.Int4Arrays != nil {
			c.SetArray(t.INT4_ARRAYS, *s.Int4Arrays)
		}
		if s.BoolArrays != nil {
			c.SetArray(t.BOOL_ARRAYS, *s.BoolArrays)
		}
		if s.CreatedAt != nil {
			c.SetTime(t.CREATED_AT, *s.CreatedAt)
		}
		if s.VJsonb != nil && s.VJsonb.Valid {
			c.SetJSON(t.V_JSONB, s.VJsonb.V)
		}
		if s.VJson != nil && s.VJson.Valid {
			c.SetJSON(t.V_JSON, s.VJson.V)
		}
		if s.VUUID != nil && s.VUUID.Valid {
			c.SetUUID(t.V_UUID, s.VUUID.V)
		}
		if s.JsonbArrays != nil && s.JsonbArrays.Valid {
			c.SetArray(t.JSONB_ARRAYS, s.JsonbArrays.V)
		}
		if s.JsonArrays != nil && s.JsonArrays.Valid {
			c.SetArray(t.JSON_ARRAYS, s.JsonArrays.V)
		}
		if s.UuidArrays != nil && s.UuidArrays.Valid {
			c.SetArray(t.UUID_ARRAYS, s.UuidArrays.V)
		}
		if s.IntAaa != nil && s.IntAaa.Valid {
			c.SetArray(t.INT_AAA, s.IntAaa.V)
		}
		if s.TsArrays != nil && s.TsArrays.Valid {
			c.SetArray(t.TS_ARRAYS, s.TsArrays.V)
		}
		if s.Int2Arrays != nil && s.Int2Arrays.Valid {
			c.SetArray(t.INT2_ARRAYS, s.Int2Arrays.V)
		}
		if s.Remark != nil && s.Remark.Valid {
			c.SetString(t.REMARK, s.Remark.V)
		}
		if s.StatusArrays != nil && s.StatusArrays.Valid {
			c.SetArray(t.STATUS_ARRAYS, s.StatusArrays.V)
		}
	}
}

func (t ARRAYS) ColumnMapper(ss ...ArraySetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t ARRAYS) RowMapperFunc() sq.RowMapper[Array] {
	return t.RowMapper
}

func (t ARRAYS) RowMapper(ctx context.Context, r *sq.Row) Array {
	v := Array{}
	v.ID = r.Int64Field(t.ID)
	var strArrays []string
	r.ArrayField(&strArrays, t.STR_ARRAYS)
	v.StrArrays = strArrays
	var int4Arrays []int32
	r.ArrayField(&int4Arrays, t.INT4_ARRAYS)
	v.Int4Arrays = int4Arrays
	var boolArrays []bool
	r.ArrayField(&boolArrays, t.BOOL_ARRAYS)
	v.BoolArrays = boolArrays
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.VJsonb = r.NullJSONField(t.V_JSONB)
	v.VJson = r.NullJSONField(t.V_JSON)
	v.VUUID = r.NullUUIDField(t.V_UUID)
	var jsonbArrays sql.Null[[]string]
	r.ArrayField(&jsonbArrays, t.JSONB_ARRAYS)
	v.JsonbArrays = jsonbArrays
	var jsonArrays sql.Null[[]string]
	r.ArrayField(&jsonArrays, t.JSON_ARRAYS)
	v.JsonArrays = jsonArrays
	var uuidArrays sql.Null[[]string]
	r.ArrayField(&uuidArrays, t.UUID_ARRAYS)
	v.UuidArrays = uuidArrays
	var intAaa sql.Null[[]int32]
	r.ArrayField(&intAaa, t.INT_AAA)
	v.IntAaa = intAaa
	var tsArrays sql.Null[[]string]
	r.ArrayField(&tsArrays, t.TS_ARRAYS)
	v.TsArrays = tsArrays
	var int2Arrays sql.Null[[]int16]
	r.ArrayField(&int2Arrays, t.INT2_ARRAYS)
	v.Int2Arrays = int2Arrays
	v.Remark = r.NullStringField(t.REMARK)
	var statusArrays sql.Null[[]any]
	r.ArrayField(&statusArrays, t.STATUS_ARRAYS)
	v.StatusArrays = statusArrays
	return v
}
