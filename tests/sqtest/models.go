package sqtest

import (
	"context"
	"database/sql"
	"time"

	"github.com/blink-io/sq"
)

type TblBasic struct {
	ID         int64                       ` db:"-" json:"-"`
	NStr       sql.Null[string]            ` db:"n_str" json:"n_str"`
	VStr       string                      ` db:"v_str" json:"v_str"`
	NEnum      sql.Null[EnumTblBasicNEnum] ` db:"n_enum" json:"n_enum"`
	VEnum      EnumTblBasicVEnum           ` db:"v_enum" json:"v_enum"`
	NInt32     sql.Null[int32]             ` db:"n_int32" json:"n_int32"`
	VInt32     int32                       ` db:"v_int32" json:"v_int32"`
	NTime      sql.Null[time.Time]         ` db:"n_time" json:"n_time"`
	VTime      time.Time                   ` db:"v_time" json:"v_time"`
	NUUID      sql.Null[[16]byte]          ` db:"n_uuid" json:"n_uuid"`
	VUUID      [16]byte                    ` db:"v_uuid" json:"v_uuid"`
	NJson      sql.Null[map[string]any]    ` db:"n_json" json:"n_json"`
	VJson      map[string]any              ` db:"v_json" json:"v_json"`
	VStrArrays []string                    ` db:"v_str_arrays" json:"v_str_arrays"`
	NStrArrays sql.Null[[]string]          ` db:"n_str_arrays" json:"n_str_arrays"`
	NBytes     sql.Null[[]byte]            ` db:"n_bytes" json:"n_bytes"`
	VBytes     []byte                      ` db:"v_bytes" json:"v_bytes"`
}

type TblBasicSetter struct {
	ID         *int64                       `db:"-" json:"-"`
	NStr       *sql.Null[string]            `db:"n_str" json:"n_str"`
	VStr       *string                      `db:"v_str" json:"v_str"`
	NEnum      *sql.Null[EnumTblBasicNEnum] `db:"n_enum" json:"n_enum"`
	VEnum      *EnumTblBasicVEnum           `db:"v_enum" json:"v_enum"`
	NInt32     *sql.Null[int32]             `db:"n_int32" json:"n_int32"`
	VInt32     *int32                       `db:"v_int32" json:"v_int32"`
	NTime      *sql.Null[time.Time]         `db:"n_time" json:"n_time"`
	VTime      *time.Time                   `db:"v_time" json:"v_time"`
	NUUID      *sql.Null[[16]byte]          `db:"n_uuid" json:"n_uuid"`
	VUUID      *[16]byte                    `db:"v_uuid" json:"v_uuid"`
	NJson      *sql.Null[map[string]any]    `db:"n_json" json:"n_json"`
	VJson      *map[string]any              `db:"v_json" json:"v_json"`
	VStrArrays *[]string                    `db:"v_str_arrays" json:"v_str_arrays"`
	NStrArrays *sql.Null[[]string]          `db:"n_str_arrays" json:"n_str_arrays"`
	NBytes     *sql.Null[[]byte]            `db:"n_bytes" json:"n_bytes"`
	VBytes     *[]byte                      `db:"v_bytes" json:"v_bytes"`
}

func (t TBL_BASIC) ColumnSetter(ctx context.Context, c *sq.Column, ss ...TblBasicSetter) {
	for idx, s := range ss {
		_ = idx
		if s.ID != nil {
			c.SetInt64(t.ID, *s.ID)
		}
		if s.NStr != nil && s.NStr.Valid {
			c.SetString(t.N_STR, s.NStr.V)
		}
		if s.VStr != nil {
			c.SetString(t.V_STR, *s.VStr)
		}
		if s.NEnum != nil && s.NEnum.Valid {
			c.SetEnum(t.N_ENUM, s.NEnum.V)
		}
		if s.VEnum != nil {
			c.SetEnum(t.V_ENUM, *s.VEnum)
		}
		if s.NInt32 != nil && s.NInt32.Valid {
			c.SetInt32(t.N_INT32, s.NInt32.V)
		}
		if s.VInt32 != nil {
			c.SetInt32(t.V_INT32, *s.VInt32)
		}
		if s.NTime != nil && s.NTime.Valid {
			c.SetTime(t.N_TIME, s.NTime.V)
		}
		if s.VTime != nil {
			c.SetTime(t.V_TIME, *s.VTime)
		}
		if s.NUUID != nil && s.NUUID.Valid {
			c.SetUUID(t.N_UUID, s.NUUID.V)
		}
		if s.VUUID != nil {
			c.SetUUID(t.V_UUID, *s.VUUID)
		}
		if s.NJson != nil && s.NJson.Valid {
			c.SetJSON(t.N_JSON, s.NJson.V)
		}
		if s.VJson != nil {
			c.SetJSON(t.V_JSON, *s.VJson)
		}
		if s.VStrArrays != nil {
			c.SetArray(t.V_STR_ARRAYS, *s.VStrArrays)
		}
		if s.NStrArrays != nil && s.NStrArrays.Valid {
			c.SetArray(t.N_STR_ARRAYS, s.NStrArrays.V)
		}
		if s.NBytes != nil && s.NBytes.Valid {
			c.SetBytes(t.N_BYTES, s.NBytes.V)
		}
		if s.VBytes != nil {
			c.SetBytes(t.V_BYTES, *s.VBytes)
		}
	}
}

func (t TBL_BASIC) ColumnMapper(ss ...TblBasicSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t TBL_BASIC) RowMapperFunc() sq.RowMapper[TblBasic] {
	return t.RowMapper
}

func (t TBL_BASIC) RowMapper(ctx context.Context, r *sq.Row) TblBasic {
	v := TblBasic{}
	v.ID = r.Int64Field(t.ID)
	v.NStr = r.NullStringField(t.N_STR)
	v.VStr = r.StringField(t.V_STR)
	var nEnum EnumTblBasicNEnum
	var nEnumValid bool
	r.NullEnumField(&nEnum, &nEnumValid, t.N_ENUM)
	v.NEnum = sql.Null[EnumTblBasicNEnum]{V: nEnum, Valid: nEnumValid}
	var vEnum EnumTblBasicVEnum
	r.EnumField(&vEnum, t.V_ENUM)
	v.VEnum = vEnum
	v.NInt32 = r.NullInt32Field(t.N_INT32)
	v.VInt32 = r.Int32Field(t.V_INT32)
	v.NTime = r.NullTimeField(t.N_TIME)
	v.VTime = r.TimeField(t.V_TIME)
	v.NUUID = r.NullUUIDField(t.N_UUID)
	v.VUUID = r.UUIDField(t.V_UUID)
	v.NJson = r.NullJSONField(t.N_JSON)
	v.VJson = r.JSONField(t.V_JSON)
	v.VStrArrays = sq.ArrayFieldFrom[[]string](r, t.V_STR_ARRAYS)
	v.NStrArrays = sq.NullArrayFieldFrom[[]string](r, t.N_STR_ARRAYS)
	v.NBytes = r.NullBytesField(t.N_BYTES)
	v.VBytes = r.BytesField(t.V_BYTES)
	return v
}
