package orm

import "context"
import "time"

import "github.com/blink-io/opt/null"
import "github.com/blink-io/opt/omitnull"
import "github.com/blink-io/sq"
import "github.com/blink-io/opt/omit"

type Array struct {
	ID          int64                    `db:"-" json:"-"`
	StrArrays   []any                    `db:"str_arrays" json:"str_arrays"`
	Int4Arrays  []any                    `db:"int4_arrays" json:"int4_arrays"`
	BoolArrays  []any                    `db:"bool_arrays" json:"bool_arrays"`
	CreatedAt   time.Time                `db:"created_at" json:"created_at"`
	VJsonb      null.Val[map[string]any] `db:"v_jsonb" json:"v_jsonb"`
	VJson       null.Val[map[string]any] `db:"v_json" json:"v_json"`
	VUUID       null.Val[[16]byte]       `db:"v_uuid" json:"v_uuid"`
	JsonbArrays null.Val[[]any]          `db:"jsonb_arrays" json:"jsonb_arrays"`
	JsonArrays  null.Val[[]any]          `db:"json_arrays" json:"json_arrays"`
	UuidArrays  null.Val[[]any]          `db:"uuid_arrays" json:"uuid_arrays"`
}

type ArraySetter struct {
	ID          omit.Val[int64]              `db:"-" json:"-"`
	StrArrays   omit.Val[[]any]              `db:"str_arrays" json:"str_arrays"`
	Int4Arrays  omit.Val[[]any]              `db:"int4_arrays" json:"int4_arrays"`
	BoolArrays  omit.Val[[]any]              `db:"bool_arrays" json:"bool_arrays"`
	CreatedAt   omit.Val[time.Time]          `db:"created_at" json:"created_at"`
	VJsonb      omitnull.Val[map[string]any] `db:"v_jsonb" json:"v_jsonb"`
	VJson       omitnull.Val[map[string]any] `db:"v_json" json:"v_json"`
	VUUID       omitnull.Val[[16]byte]       `db:"v_uuid" json:"v_uuid"`
	JsonbArrays omitnull.Val[[]any]          `db:"jsonb_arrays" json:"jsonb_arrays"`
	JsonArrays  omitnull.Val[[]any]          `db:"json_arrays" json:"json_arrays"`
	UuidArrays  omitnull.Val[[]any]          `db:"uuid_arrays" json:"uuid_arrays"`
}

func (t ARRAYS) ColumnMapper(ctx context.Context, c *sq.Column, s ArraySetter) {
	s.ID.IfSet(func(v int64) {
		c.SetInt64(t.ID, v)
	})
	s.StrArrays.IfSet(func(v []any) {
		c.SetArray(t.STR_ARRAYS, v)
	})
	s.Int4Arrays.IfSet(func(v []any) {
		c.SetArray(t.INT4_ARRAYS, v)
	})
	s.BoolArrays.IfSet(func(v []any) {
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
	s.JsonbArrays.IfSet(func(v []any) {
		c.SetArray(t.JSONB_ARRAYS, v)
	})
	s.JsonArrays.IfSet(func(v []any) {
		c.SetArray(t.JSON_ARRAYS, v)
	})
	s.UuidArrays.IfSet(func(v []any) {
		c.SetArray(t.UUID_ARRAYS, v)
	})
}

func (t ARRAYS) RowMapper(ctx context.Context, r *sq.Row) Array {
	v := Array{}
	v.ID = r.Int64Field(t.ID)
	var strArrays []any
	r.ScanField(strArrays, t.STR_ARRAYS)
	v.StrArrays = strArrays
	var int4Arrays []any
	r.ScanField(int4Arrays, t.INT4_ARRAYS)
	v.Int4Arrays = int4Arrays
	var boolArrays []any
	r.ScanField(boolArrays, t.BOOL_ARRAYS)
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
	var jsonbArrays = new([]any)
	r.ScanField(jsonbArrays, t.JSONB_ARRAYS)
	v.JsonbArrays = null.FromPtr(jsonbArrays)
	var jsonArrays = new([]any)
	r.ScanField(jsonArrays, t.JSON_ARRAYS)
	v.JsonArrays = null.FromPtr(jsonArrays)
	var uuidArrays = new([]any)
	r.ScanField(uuidArrays, t.UUID_ARRAYS)
	v.UuidArrays = null.FromPtr(uuidArrays)
	return v
}

type Device struct {
	ID        int64     `db:"-" json:"-"`
	Name      string    `db:"name" json:"name"`
	Model     string    `db:"model" json:"model"`
	GUID      string    `db:"guid" json:"guid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type DeviceSetter struct {
	ID        omit.Val[int64]     `db:"-" json:"-"`
	Name      omit.Val[string]    `db:"name" json:"name"`
	Model     omit.Val[string]    `db:"model" json:"model"`
	GUID      omit.Val[string]    `db:"guid" json:"guid"`
	CreatedAt omit.Val[time.Time] `db:"created_at" json:"created_at"`
	UpdatedAt omit.Val[time.Time] `db:"updated_at" json:"updated_at"`
}

func (t DEVICES) ColumnMapper(ctx context.Context, c *sq.Column, s DeviceSetter) {
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
	ID        int64                    `db:"-" json:"-"`
	Status    EnumEnumsStatus          `db:"status" json:"status"`
	CreatedAt time.Time                `db:"created_at" json:"created_at"`
	Moodx     null.Val[EnumEnumsMoodx] `db:"moodx" json:"moodx"`
}

type EnumSetter struct {
	ID        omit.Val[int64]              `db:"-" json:"-"`
	Status    omit.Val[EnumEnumsStatus]    `db:"status" json:"status"`
	CreatedAt omit.Val[time.Time]          `db:"created_at" json:"created_at"`
	Moodx     omitnull.Val[EnumEnumsMoodx] `db:"moodx" json:"moodx"`
}

func (t ENUMS) ColumnMapper(ctx context.Context, c *sq.Column, s EnumSetter) {
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
	ID int64 `db:"-" json:"-"`
}

type HelloWorldSetter struct {
	ID omit.Val[int64] `db:"-" json:"-"`
}

func (t HELLO_WORLD) ColumnMapper(ctx context.Context, c *sq.Column, s HelloWorldSetter) {
	s.ID.IfSet(func(v int64) {
		c.SetInt64(t.ID, v)
	})
}

func (t HELLO_WORLD) RowMapper(ctx context.Context, r *sq.Row) HelloWorld {
	v := HelloWorld{}
	v.ID = r.Int64Field(t.ID)
	return v
}

type Log struct {
	ID        int64               `db:"-" json:"-"`
	Content   string              `db:"content" json:"content"`
	CreatedAt time.Time           `db:"created_at" json:"created_at"`
	AuditedAt null.Val[time.Time] `db:"audited_at" json:"audited_at"`
}

type LogSetter struct {
	ID        omit.Val[int64]         `db:"-" json:"-"`
	Content   omit.Val[string]        `db:"content" json:"content"`
	CreatedAt omit.Val[time.Time]     `db:"created_at" json:"created_at"`
	AuditedAt omitnull.Val[time.Time] `db:"audited_at" json:"audited_at"`
}

func (t LOGS) ColumnMapper(ctx context.Context, c *sq.Column, s LogSetter) {
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
	Id1       int32     `db:"id1" json:"id1"`
	Id2       int32     `db:"id2" json:"id2"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	GUID      string    `db:"guid" json:"guid"`
}

type MkeySetter struct {
	Id1       omit.Val[int32]     `db:"id1" json:"id1"`
	Id2       omit.Val[int32]     `db:"id2" json:"id2"`
	Name      omit.Val[string]    `db:"name" json:"name"`
	CreatedAt omit.Val[time.Time] `db:"created_at" json:"created_at"`
	GUID      omit.Val[string]    `db:"guid" json:"guid"`
}

func (t MKEYS) ColumnMapper(ctx context.Context, c *sq.Column, s MkeySetter) {
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
	ID        int64     `db:"-" json:"-"`
	GUID      string    `db:"guid" json:"guid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Content   string    `db:"content" json:"content"`
}

type NewWordSetter struct {
	ID        omit.Val[int64]     `db:"-" json:"-"`
	GUID      omit.Val[string]    `db:"guid" json:"guid"`
	CreatedAt omit.Val[time.Time] `db:"created_at" json:"created_at"`
	UpdatedAt omit.Val[time.Time] `db:"updated_at" json:"updated_at"`
	Content   omit.Val[string]    `db:"content" json:"content"`
}

func (t NEW_WORDS) ColumnMapper(ctx context.Context, c *sq.Column, s NewWordSetter) {
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
	ID          int64            `db:"-" json:"-"`
	Name        string           `db:"name" json:"name"`
	Code        string           `db:"code" json:"code"`
	Description null.Val[string] `db:"description" json:"description"`
	GUID        string           `db:"guid" json:"guid"`
	CreatedAt   time.Time        `db:"created_at" json:"created_at"`
}

type TagSetter struct {
	ID          omit.Val[int64]      `db:"-" json:"-"`
	Name        omit.Val[string]     `db:"name" json:"name"`
	Code        omit.Val[string]     `db:"code" json:"code"`
	Description omitnull.Val[string] `db:"description" json:"description"`
	GUID        omit.Val[string]     `db:"guid" json:"guid"`
	CreatedAt   omit.Val[time.Time]  `db:"created_at" json:"created_at"`
}

func (t TAGS) ColumnMapper(ctx context.Context, c *sq.Column, s TagSetter) {
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
	ID          int64            `db:"-" json:"-"`
	Name        string           `db:"name" json:"name"`
	Code        string           `db:"code" json:"code"`
	Description null.Val[string] `db:"description" json:"description"`
	GUID        string           `db:"guid" json:"guid"`
	CreatedAt   time.Time        `db:"created_at" json:"created_at"`
}

type TagsBakSetter struct {
	ID          omit.Val[int64]      `db:"-" json:"-"`
	Name        omit.Val[string]     `db:"name" json:"name"`
	Code        omit.Val[string]     `db:"code" json:"code"`
	Description omitnull.Val[string] `db:"description" json:"description"`
	GUID        omit.Val[string]     `db:"guid" json:"guid"`
	CreatedAt   omit.Val[time.Time]  `db:"created_at" json:"created_at"`
}

func (t TAGS_BAK) ColumnMapper(ctx context.Context, c *sq.Column, s TagsBakSetter) {
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
	ID          int64            `db:"-" json:"-"`
	UserID      int64            `db:"user_id" json:"user_id"`
	GUID        string           `db:"guid" json:"guid"`
	Model       string           `db:"model" json:"model"`
	Name        string           `db:"name" json:"name"`
	Description null.Val[string] `db:"description" json:"description"`
	CreatedAt   time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at" json:"updated_at"`
}

type UserDeviceSetter struct {
	ID          omit.Val[int64]      `db:"-" json:"-"`
	UserID      omit.Val[int64]      `db:"user_id" json:"user_id"`
	GUID        omit.Val[string]     `db:"guid" json:"guid"`
	Model       omit.Val[string]     `db:"model" json:"model"`
	Name        omit.Val[string]     `db:"name" json:"name"`
	Description omitnull.Val[string] `db:"description" json:"description"`
	CreatedAt   omit.Val[time.Time]  `db:"created_at" json:"created_at"`
	UpdatedAt   omit.Val[time.Time]  `db:"updated_at" json:"updated_at"`
}

func (t USER_DEVICES) ColumnMapper(ctx context.Context, c *sq.Column, s UserDeviceSetter) {
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
	ID        int64     `db:"-" json:"-"`
	Username  string    `db:"username" json:"username"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Level     int16     `db:"level" json:"level"`
	Score     float64   `db:"score" json:"score"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	GUID      string    `db:"guid" json:"guid"`
	TenantID  int64     `db:"tenant_id" json:"tenant_id"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserSetter struct {
	ID        omit.Val[int64]     `db:"-" json:"-"`
	Username  omit.Val[string]    `db:"username" json:"username"`
	FirstName omit.Val[string]    `db:"first_name" json:"first_name"`
	LastName  omit.Val[string]    `db:"last_name" json:"last_name"`
	Level     omit.Val[int16]     `db:"level" json:"level"`
	Score     omit.Val[float64]   `db:"score" json:"score"`
	CreatedAt omit.Val[time.Time] `db:"created_at" json:"created_at"`
	GUID      omit.Val[string]    `db:"guid" json:"guid"`
	TenantID  omit.Val[int64]     `db:"tenant_id" json:"tenant_id"`
	UpdatedAt omit.Val[time.Time] `db:"updated_at" json:"updated_at"`
}

func (t USERS) ColumnMapper(ctx context.Context, c *sq.Column, s UserSetter) {
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
