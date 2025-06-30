package orm

import (
	"context"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/blink-io/sq"
)

type Array struct {
	ID          int64                    ` db:"-" json:"-"`
	StrArrays   []string                 ` db:"str_arrays" json:"str_arrays"`
	Int4Arrays  []int32                  ` db:"int4_arrays" json:"int4_arrays"`
	Int2Arrays  []int16                  ` db:"int2_arrays" json:"int2_arrays"`
	BoolArrays  []bool                   ` db:"bool_arrays" json:"bool_arrays"`
	CreatedAt   time.Time                ` db:"created_at" json:"created_at"`
	VJsonb      null.Val[map[string]any] ` db:"v_jsonb" json:"v_jsonb"`
	VJson       null.Val[map[string]any] ` db:"v_json" json:"v_json"`
	VUUID       null.Val[[16]byte]       ` db:"v_uuid" json:"v_uuid"`
	JsonbArrays null.Val[[]string]       ` db:"jsonb_arrays" json:"jsonb_arrays"`
	JsonArrays  null.Val[[]string]       ` db:"json_arrays" json:"json_arrays"`
	UuidArrays  null.Val[[]string]       ` db:"uuid_arrays" json:"uuid_arrays"`
	IntAaa      null.Val[[]int32]        ` db:"int_aaa" json:"int_aaa"`
	TsArrays    null.Val[[]time.Time]    ` db:"ts_arrays" json:"ts_arrays"`
}

type ArraySetter struct {
	ID          omit.Val[int64]              ` db:"-" json:"-"`
	StrArrays   omit.Val[[]string]           ` db:"str_arrays" json:"str_arrays"`
	Int4Arrays  omit.Val[[]int32]            ` db:"int4_arrays" json:"int4_arrays"`
	Int2Arrays  omitnull.Val[[]int16]        ` db:"int4_arrays" json:"int2_arrays"`
	BoolArrays  omit.Val[[]bool]             ` db:"bool_arrays" json:"bool_arrays"`
	CreatedAt   omit.Val[time.Time]          ` db:"created_at" json:"created_at"`
	VJsonb      omitnull.Val[map[string]any] ` db:"v_jsonb" json:"v_jsonb"`
	VJson       omitnull.Val[map[string]any] ` db:"v_json" json:"v_json"`
	VUUID       omitnull.Val[[16]byte]       ` db:"v_uuid" json:"v_uuid"`
	JsonbArrays omitnull.Val[[]string]       ` db:"jsonb_arrays" json:"jsonb_arrays"`
	JsonArrays  omitnull.Val[[]string]       ` db:"json_arrays" json:"json_arrays"`
	UuidArrays  omitnull.Val[[]string]       ` db:"uuid_arrays" json:"uuid_arrays"`
	IntAaa      omitnull.Val[[]int32]        ` db:"int_aaa" json:"int_aaa"`
	TsArrays    omitnull.Val[[]time.Time]    ` db:"ts_arrays" json:"ts_arrays"`
}

func (t ARRAYS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...ArraySetter) {
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
		s.Int2Arrays.IfSet(func(v []int16) {
			c.SetArray(t.INT2_ARRAYS, v)
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
		s.JsonbArrays.IfSet(func(v []string) {
			c.SetArray(t.JSONB_ARRAYS, v)
		})
		s.JsonArrays.IfSet(func(v []string) {
			c.SetArray(t.JSON_ARRAYS, v)
		})
		s.UuidArrays.IfSet(func(v []string) {
			c.SetArray(t.UUID_ARRAYS, v)
		})
		s.IntAaa.IfSet(func(v []int32) {
			c.SetArray(t.INT_AAA, v)
		})
		s.TsArrays.IfSet(func(v []time.Time) {
			c.SetArray(t.TS_ARRAYS, v)
		})
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
	r.ArrayField(strArrays, t.STR_ARRAYS)
	v.StrArrays = strArrays
	var int4Arrays []int32
	r.ArrayField(int4Arrays, t.INT4_ARRAYS)
	v.Int4Arrays = int4Arrays
	var boolArrays []bool
	r.ArrayField(boolArrays, t.BOOL_ARRAYS)
	v.BoolArrays = boolArrays
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	var vJsonb = new(map[string]any)
	r.JSONField(vJsonb, t.V_JSONB)
	v.VJsonb = null.FromPtr(vJsonb)
	var vJson = new(map[string]any)
	r.JSONField(vJson, t.V_JSON)
	v.VJson = null.FromPtr(vJson)
	var vUUID = new([16]byte)
	r.UUIDField(vUUID, t.V_UUID)
	v.VUUID = null.FromPtr(vUUID)
	var jsonbArrays = new([]string)
	r.ArrayField(jsonbArrays, t.JSONB_ARRAYS)
	v.JsonbArrays = null.FromPtr(jsonbArrays)
	var jsonArrays = new([]string)
	r.ArrayField(jsonArrays, t.JSON_ARRAYS)
	v.JsonArrays = null.FromPtr(jsonArrays)
	var uuidArrays = new([]string)
	r.ArrayField(uuidArrays, t.UUID_ARRAYS)
	v.UuidArrays = null.FromPtr(uuidArrays)
	var intAaa = new([]int32)
	r.ArrayField(intAaa, t.INT_AAA)
	v.IntAaa = null.FromPtr(intAaa)
	var tsArrays = new([]time.Time)
	r.ArrayField(tsArrays, t.TS_ARRAYS)
	v.TsArrays = null.FromPtr(tsArrays)
	return v
}

type Device struct {
	ID        int64     ` db:"-" json:"-"`
	Name      string    ` db:"name" json:"name"`
	Model     string    ` db:"model" json:"model"`
	GUID      string    ` db:"guid" json:"guid"`
	CreatedAt time.Time ` db:"created_at" json:"created_at"`
	UpdatedAt time.Time ` db:"updated_at" json:"updated_at"`
}

type DeviceSetter struct {
	ID        omit.Val[int64]     ` db:"-" json:"-"`
	Name      omit.Val[string]    ` db:"name" json:"name"`
	Model     omit.Val[string]    ` db:"model" json:"model"`
	GUID      omit.Val[string]    ` db:"guid" json:"guid"`
	CreatedAt omit.Val[time.Time] ` db:"created_at" json:"created_at"`
	UpdatedAt omit.Val[time.Time] ` db:"updated_at" json:"updated_at"`
}

func (t DEVICES) ColumnSetter(ctx context.Context, c *sq.Column, ss ...DeviceSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Name.IfSet(func(v string) {
			c.SetString(t.NAME, v)
		})
		s.Model.IfSet(func(v string) {
			c.SetString(t.MODEL, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.UpdatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.UPDATED_AT, v)
		})
	}
}

func (t DEVICES) ColumnMapper(ss ...DeviceSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t DEVICES) RowMapperFunc() sq.RowMapper[Device] {
	return t.RowMapper
}

func (t DEVICES) RowMapper(ctx context.Context, r *sq.Row) Device {
	v := Device{}
	v.ID = r.Int64Field(t.ID)
	v.Name = r.StringField(t.NAME)
	v.Model = r.StringField(t.MODEL)
	v.GUID = r.StringField(t.GUID)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.UpdatedAt = r.TimeField(t.UPDATED_AT)
	return v
}

type Enum struct {
	ID        int64                    ` db:"-" json:"-"`
	Status    EnumEnumsStatus          ` db:"status" json:"status"`
	CreatedAt time.Time                ` db:"created_at" json:"created_at"`
	Moodx     null.Val[EnumEnumsMoodx] ` db:"moodx" json:"moodx"`
}

type EnumSetter struct {
	ID        omit.Val[int64]              ` db:"-" json:"-"`
	Status    omit.Val[EnumEnumsStatus]    ` db:"status" json:"status"`
	CreatedAt omit.Val[time.Time]          ` db:"created_at" json:"created_at"`
	Moodx     omitnull.Val[EnumEnumsMoodx] ` db:"moodx" json:"moodx"`
}

func (t ENUMS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...EnumSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Status.IfSet(func(v EnumEnumsStatus) {
			c.SetEnum(t.STATUS, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.Moodx.IfSet(func(v EnumEnumsMoodx) {
			c.SetEnum(t.MOODX, v)
		})
	}
}

func (t ENUMS) ColumnMapper(ss ...EnumSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t ENUMS) RowMapperFunc() sq.RowMapper[Enum] {
	return t.RowMapper
}

func (t ENUMS) RowMapper(ctx context.Context, r *sq.Row) Enum {
	v := Enum{}
	v.ID = r.Int64Field(t.ID)
	var status EnumEnumsStatus
	r.EnumField(status, t.STATUS)
	v.Status = status
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	var moodx = new(EnumEnumsMoodx)
	r.EnumField(moodx, t.MOODX)
	v.Moodx = null.FromPtr(moodx)
	return v
}

type HelloWorld struct {
	ID int64 ` db:"-" json:"-"`
}

type HelloWorldSetter struct {
	ID omit.Val[int64] ` db:"-" json:"-"`
}

func (t HELLO_WORLD) ColumnSetter(ctx context.Context, c *sq.Column, ss ...HelloWorldSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
	}
}

func (t HELLO_WORLD) ColumnMapper(ss ...HelloWorldSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t HELLO_WORLD) RowMapperFunc() sq.RowMapper[HelloWorld] {
	return t.RowMapper
}

func (t HELLO_WORLD) RowMapper(ctx context.Context, r *sq.Row) HelloWorld {
	v := HelloWorld{}
	v.ID = r.Int64Field(t.ID)
	return v
}

type Log struct {
	ID        int64               ` db:"-" json:"-"`
	Content   string              ` db:"content" json:"content"`
	CreatedAt time.Time           ` db:"created_at" json:"created_at"`
	AuditedAt null.Val[time.Time] ` db:"audited_at" json:"audited_at"`
}

type LogSetter struct {
	ID        omit.Val[int64]         ` db:"-" json:"-"`
	Content   omit.Val[string]        ` db:"content" json:"content"`
	CreatedAt omit.Val[time.Time]     ` db:"created_at" json:"created_at"`
	AuditedAt omitnull.Val[time.Time] ` db:"audited_at" json:"audited_at"`
}

func (t LOGS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...LogSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Content.IfSet(func(v string) {
			c.SetString(t.CONTENT, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.AuditedAt.IfSet(func(v time.Time) {
			c.SetTime(t.AUDITED_AT, v)
		})
	}
}

func (t LOGS) ColumnMapper(ss ...LogSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t LOGS) RowMapperFunc() sq.RowMapper[Log] {
	return t.RowMapper
}

func (t LOGS) RowMapper(ctx context.Context, r *sq.Row) Log {
	v := Log{}
	v.ID = r.Int64Field(t.ID)
	v.Content = r.StringField(t.CONTENT)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	auditedAt := r.NullTimeField(t.AUDITED_AT)
	v.AuditedAt = null.FromCond(auditedAt.V, auditedAt.Valid)
	return v
}

type Mkey struct {
	Id1       int32     ` db:"id1" json:"id1"`
	Id2       int32     ` db:"id2" json:"id2"`
	Name      string    ` db:"name" json:"name"`
	CreatedAt time.Time ` db:"created_at" json:"created_at"`
	GUID      string    ` db:"guid" json:"guid"`
}

type MkeySetter struct {
	Id1       omit.Val[int32]     ` db:"id1" json:"id1"`
	Id2       omit.Val[int32]     ` db:"id2" json:"id2"`
	Name      omit.Val[string]    ` db:"name" json:"name"`
	CreatedAt omit.Val[time.Time] ` db:"created_at" json:"created_at"`
	GUID      omit.Val[string]    ` db:"guid" json:"guid"`
}

func (t MKEYS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...MkeySetter) {
	for idx, s := range ss {
		_ = idx
		s.Id1.IfSet(func(v int32) {
			c.SetInt32(t.ID1, v)
		})
		s.Id2.IfSet(func(v int32) {
			c.SetInt32(t.ID2, v)
		})
		s.Name.IfSet(func(v string) {
			c.SetString(t.NAME, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
	}
}

func (t MKEYS) ColumnMapper(ss ...MkeySetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t MKEYS) RowMapperFunc() sq.RowMapper[Mkey] {
	return t.RowMapper
}

func (t MKEYS) RowMapper(ctx context.Context, r *sq.Row) Mkey {
	v := Mkey{}
	v.Id1 = r.Int32Field(t.ID1)
	v.Id2 = r.Int32Field(t.ID2)
	v.Name = r.StringField(t.NAME)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.GUID = r.StringField(t.GUID)
	return v
}

type NewWord struct {
	ID        int64     ` db:"-" json:"-"`
	GUID      string    ` db:"guid" json:"guid"`
	CreatedAt time.Time ` db:"created_at" json:"created_at"`
	UpdatedAt time.Time ` db:"updated_at" json:"updated_at"`
	Content   string    ` db:"content" json:"content"`
}

type NewWordSetter struct {
	ID        omit.Val[int64]     ` db:"-" json:"-"`
	GUID      omit.Val[string]    ` db:"guid" json:"guid"`
	CreatedAt omit.Val[time.Time] ` db:"created_at" json:"created_at"`
	UpdatedAt omit.Val[time.Time] ` db:"updated_at" json:"updated_at"`
	Content   omit.Val[string]    ` db:"content" json:"content"`
}

func (t NEW_WORDS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...NewWordSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.UpdatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.UPDATED_AT, v)
		})
		s.Content.IfSet(func(v string) {
			c.SetString(t.CONTENT, v)
		})
	}
}

func (t NEW_WORDS) ColumnMapper(ss ...NewWordSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t NEW_WORDS) RowMapperFunc() sq.RowMapper[NewWord] {
	return t.RowMapper
}

func (t NEW_WORDS) RowMapper(ctx context.Context, r *sq.Row) NewWord {
	v := NewWord{}
	v.ID = r.Int64Field(t.ID)
	v.GUID = r.StringField(t.GUID)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.UpdatedAt = r.TimeField(t.UPDATED_AT)
	v.Content = r.StringField(t.CONTENT)
	return v
}

type Tag struct {
	ID          int64            ` db:"-" json:"-"`
	Name        string           ` db:"name" json:"name"`
	Code        string           ` db:"code" json:"code"`
	Description null.Val[string] ` db:"description" json:"description"`
	GUID        string           ` db:"guid" json:"guid"`
	CreatedAt   time.Time        ` db:"created_at" json:"created_at"`
}

type TagSetter struct {
	ID          omit.Val[int64]      ` db:"-" json:"-"`
	Name        omit.Val[string]     ` db:"name" json:"name"`
	Code        omit.Val[string]     ` db:"code" json:"code"`
	Description omitnull.Val[string] ` db:"description" json:"description"`
	GUID        omit.Val[string]     ` db:"guid" json:"guid"`
	CreatedAt   omit.Val[time.Time]  ` db:"created_at" json:"created_at"`
}

func (t TAGS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...TagSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Name.IfSet(func(v string) {
			c.SetString(t.NAME, v)
		})
		s.Code.IfSet(func(v string) {
			c.SetString(t.CODE, v)
		})
		s.Description.IfSet(func(v string) {
			c.SetString(t.DESCRIPTION, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
	}
}

func (t TAGS) ColumnMapper(ss ...TagSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t TAGS) RowMapperFunc() sq.RowMapper[Tag] {
	return t.RowMapper
}

func (t TAGS) RowMapper(ctx context.Context, r *sq.Row) Tag {
	v := Tag{}
	v.ID = r.Int64Field(t.ID)
	v.Name = r.StringField(t.NAME)
	v.Code = r.StringField(t.CODE)
	description := r.NullStringField(t.DESCRIPTION)
	v.Description = null.FromCond(description.V, description.Valid)
	v.GUID = r.StringField(t.GUID)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	return v
}

type TagsBak struct {
	ID          int64            ` db:"-" json:"-"`
	Name        string           ` db:"name" json:"name"`
	Code        string           ` db:"code" json:"code"`
	Description null.Val[string] ` db:"description" json:"description"`
	GUID        string           ` db:"guid" json:"guid"`
	CreatedAt   time.Time        ` db:"created_at" json:"created_at"`
}

type TagsBakSetter struct {
	ID          omit.Val[int64]      ` db:"-" json:"-"`
	Name        omit.Val[string]     ` db:"name" json:"name"`
	Code        omit.Val[string]     ` db:"code" json:"code"`
	Description omitnull.Val[string] ` db:"description" json:"description"`
	GUID        omit.Val[string]     ` db:"guid" json:"guid"`
	CreatedAt   omit.Val[time.Time]  ` db:"created_at" json:"created_at"`
}

func (t TAGS_BAK) ColumnSetter(ctx context.Context, c *sq.Column, ss ...TagsBakSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Name.IfSet(func(v string) {
			c.SetString(t.NAME, v)
		})
		s.Code.IfSet(func(v string) {
			c.SetString(t.CODE, v)
		})
		s.Description.IfSet(func(v string) {
			c.SetString(t.DESCRIPTION, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
	}
}

func (t TAGS_BAK) ColumnMapper(ss ...TagsBakSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t TAGS_BAK) RowMapperFunc() sq.RowMapper[TagsBak] {
	return t.RowMapper
}

func (t TAGS_BAK) RowMapper(ctx context.Context, r *sq.Row) TagsBak {
	v := TagsBak{}
	v.ID = r.Int64Field(t.ID)
	v.Name = r.StringField(t.NAME)
	v.Code = r.StringField(t.CODE)
	description := r.NullStringField(t.DESCRIPTION)
	v.Description = null.FromCond(description.V, description.Valid)
	v.GUID = r.StringField(t.GUID)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	return v
}

type UserDevice struct {
	ID          int64            ` db:"-" json:"-"`
	UserID      int64            ` db:"user_id" json:"user_id"`
	GUID        string           ` db:"guid" json:"guid"`
	Model       string           ` db:"model" json:"model"`
	Name        string           ` db:"name" json:"name"`
	Description null.Val[string] ` db:"description" json:"description"`
	CreatedAt   time.Time        ` db:"created_at" json:"created_at"`
	UpdatedAt   time.Time        ` db:"updated_at" json:"updated_at"`
}

type UserDeviceSetter struct {
	ID          omit.Val[int64]      ` db:"-" json:"-"`
	UserID      omit.Val[int64]      ` db:"user_id" json:"user_id"`
	GUID        omit.Val[string]     ` db:"guid" json:"guid"`
	Model       omit.Val[string]     ` db:"model" json:"model"`
	Name        omit.Val[string]     ` db:"name" json:"name"`
	Description omitnull.Val[string] ` db:"description" json:"description"`
	CreatedAt   omit.Val[time.Time]  ` db:"created_at" json:"created_at"`
	UpdatedAt   omit.Val[time.Time]  ` db:"updated_at" json:"updated_at"`
}

func (t USER_DEVICES) ColumnSetter(ctx context.Context, c *sq.Column, ss ...UserDeviceSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.UserID.IfSet(func(v int64) {
			c.SetInt64(t.USER_ID, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.Model.IfSet(func(v string) {
			c.SetString(t.MODEL, v)
		})
		s.Name.IfSet(func(v string) {
			c.SetString(t.NAME, v)
		})
		s.Description.IfSet(func(v string) {
			c.SetString(t.DESCRIPTION, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.UpdatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.UPDATED_AT, v)
		})
	}
}

func (t USER_DEVICES) ColumnMapper(ss ...UserDeviceSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t USER_DEVICES) RowMapperFunc() sq.RowMapper[UserDevice] {
	return t.RowMapper
}

func (t USER_DEVICES) RowMapper(ctx context.Context, r *sq.Row) UserDevice {
	v := UserDevice{}
	v.ID = r.Int64Field(t.ID)
	v.UserID = r.Int64Field(t.USER_ID)
	v.GUID = r.StringField(t.GUID)
	v.Model = r.StringField(t.MODEL)
	v.Name = r.StringField(t.NAME)
	description := r.NullStringField(t.DESCRIPTION)
	v.Description = null.FromCond(description.V, description.Valid)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.UpdatedAt = r.TimeField(t.UPDATED_AT)
	return v
}

type User struct {
	ID        int64     ` db:"-" json:"-"`
	Username  string    ` db:"username" json:"username"`
	FirstName string    ` db:"first_name" json:"first_name"`
	LastName  string    ` db:"last_name" json:"last_name"`
	Level     int16     ` db:"level" json:"level"`
	Score     float64   ` db:"score" json:"score"`
	CreatedAt time.Time ` db:"created_at" json:"created_at"`
	GUID      string    ` db:"guid" json:"guid"`
	TenantID  int64     ` db:"tenant_id" json:"tenant_id"`
	UpdatedAt time.Time ` db:"updated_at" json:"updated_at"`
}

type UserSetter struct {
	ID        omit.Val[int64]     ` db:"-" json:"-"`
	Username  omit.Val[string]    ` db:"username" json:"username"`
	FirstName omit.Val[string]    ` db:"first_name" json:"first_name"`
	LastName  omit.Val[string]    ` db:"last_name" json:"last_name"`
	Level     omit.Val[int16]     ` db:"level" json:"level"`
	Score     omit.Val[float64]   ` db:"score" json:"score"`
	CreatedAt omit.Val[time.Time] ` db:"created_at" json:"created_at"`
	GUID      omit.Val[string]    ` db:"guid" json:"guid"`
	TenantID  omit.Val[int64]     ` db:"tenant_id" json:"tenant_id"`
	UpdatedAt omit.Val[time.Time] ` db:"updated_at" json:"updated_at"`
}

func (t USERS) ColumnSetter(ctx context.Context, c *sq.Column, ss ...UserSetter) {
	for idx, s := range ss {
		_ = idx
		s.ID.IfSet(func(v int64) {
			c.SetInt64(t.ID, v)
		})
		s.Username.IfSet(func(v string) {
			c.SetString(t.USERNAME, v)
		})
		s.FirstName.IfSet(func(v string) {
			c.SetString(t.FIRST_NAME, v)
		})
		s.LastName.IfSet(func(v string) {
			c.SetString(t.LAST_NAME, v)
		})
		s.Level.IfSet(func(v int16) {
			c.SetInt16(t.LEVEL, v)
		})
		s.Score.IfSet(func(v float64) {
			c.SetFloat64(t.SCORE, v)
		})
		s.CreatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.CREATED_AT, v)
		})
		s.GUID.IfSet(func(v string) {
			c.SetString(t.GUID, v)
		})
		s.TenantID.IfSet(func(v int64) {
			c.SetInt64(t.TENANT_ID, v)
		})
		s.UpdatedAt.IfSet(func(v time.Time) {
			c.SetTime(t.UPDATED_AT, v)
		})
	}
}

func (t USERS) ColumnMapper(ss ...UserSetter) sq.ColumnMapper {
	return func(ctx context.Context, c *sq.Column) {
		t.ColumnSetter(ctx, c, ss...)
	}
}

func (t USERS) RowMapperFunc() sq.RowMapper[User] {
	return t.RowMapper
}

func (t USERS) RowMapper(ctx context.Context, r *sq.Row) User {
	v := User{}
	v.ID = r.Int64Field(t.ID)
	v.Username = r.StringField(t.USERNAME)
	v.FirstName = r.StringField(t.FIRST_NAME)
	v.LastName = r.StringField(t.LAST_NAME)
	v.Level = r.Int16Field(t.LEVEL)
	v.Score = r.Float64Field(t.SCORE)
	v.CreatedAt = r.TimeField(t.CREATED_AT)
	v.GUID = r.StringField(t.GUID)
	v.TenantID = r.Int64Field(t.TENANT_ID)
	v.UpdatedAt = r.TimeField(t.UPDATED_AT)
	return v
}
