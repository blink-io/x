package buntest

import (
	"database/sql"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"time"

	"github.com/blink-io/hyperbun/schema"
	"github.com/blink-io/sq"
	"github.com/blink-io/x/ptr"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

type EnumTblBasicNEnum string
type EnumTblBasicVEnum string

func randomTblDemoSetter() *TblDemoSetter {
	vuuid := uuid.NewString()
	vstr := gofakeit.Name()
	venum := EnumTblBasicVEnum("active")
	nenum := EnumTblBasicNEnum("active")
	vjson := map[string]any{"v": gofakeit.City()}

	var s = &TblDemoSetter{
		NUUID:  ptr.Of(sq.ValidFrom(uuid.NewString())),
		NEnum:  ptr.Of(sql.Null[EnumTblBasicNEnum]{V: nenum, Valid: true}),
		NStr:   ptr.Of(sq.NullString{}),
		NTime:  ptr.Of(sq.NullTime{}),
		NInt32: ptr.Of(sq.ValidFrom(gofakeit.Int32())),

		VUUID:      ptr.Of(vuuid),
		VStr:       &vstr,
		VJson:      &vjson,
		VInt32:     ptr.Of(gofakeit.Int32()),
		VEnum:      ptr.Of(venum),
		VTime:      ptr.Of(gofakeit.Date()),
		VStrArrays: ptr.Of(pq.StringArray{gofakeit.Name(), gofakeit.City(), gofakeit.AppName()}),
		VBytes:     ptr.Of(pq.ByteaArray{[]byte(gofakeit.Name())}),
	}

	return s
}

func randomTblDemo() *TblDemo {
	vuuid := uuid.NewString()
	vstr := gofakeit.Name()
	venum := EnumTblBasicVEnum("active")
	nenum := EnumTblBasicNEnum("active")
	vjson := map[string]any{"v": gofakeit.City()}

	var s = &TblDemo{
		NUUID:        ptr.Of(uuid.NewString()),
		NEnum:        ptr.Of(nenum),
		NStr:         nil,
		NTime:        nil,
		NInt32:       ptr.Of(gofakeit.Int32()),
		NInt32Arrays: ptr.Of([]int32{gofakeit.Int32(), gofakeit.Int32()}),

		VUUID:      vuuid,
		VStr:       vstr,
		VJson:      vjson,
		VInt32:     gofakeit.Int32(),
		VEnum:      venum,
		VTime:      gofakeit.Date(),
		VStrArrays: []string{gofakeit.Name(), gofakeit.City(), gofakeit.AppName()},
		VBytes:     []byte(gofakeit.Name()),
	}

	if int(gofakeit.Int32()/2) == 0 {
		s.ID = gofakeit.Int64()
	}

	return s
}

type TblDemo struct {
	bun.BaseModel `bun:"table:tbl_demo,alias:u2"`
	ID            int64              `db:"-" bun:"id,pk,nullzero"`
	NStr          *string            `db:"n_str" bun:"n_str"`
	VStr          string             `db:"v_str" bun:"v_str"`
	NEnum         *EnumTblBasicNEnum `db:"n_enum" bun:"n_enum"`
	VEnum         EnumTblBasicVEnum  `db:"v_enum" bun:"v_enum"`
	NInt32        *int32             `db:"n_int32" bun:"n_int32"`
	VInt32        int32              `db:"v_int32" bun:"v_int32"`
	NTime         *time.Time         `db:"n_time" bun:"n_time"`
	VTime         time.Time          `db:"v_time" bun:"v_time"`
	NUUID         *string            `db:"n_uuid" bun:"n_uuid"`
	VUUID         string             `db:"v_uuid" bun:"v_uuid"`
	NJson         map[string]any     `db:"n_json" bun:"n_json"`
	VJson         map[string]any     `db:"v_json" bun:"v_json,json_use_number"`
	VStrArrays    []string           `db:"v_str_arrays" bun:"v_str_arrays,array"`
	NStrArrays    *[]string          `db:"n_str_arrays" bun:"n_str_arrays,array"`
	NBytes        *[]byte            `db:"n_bytes" bun:"n_bytes"`
	VBytes        []byte             `db:"v_bytes" bun:"v_bytes"`
	NInt32Arrays  *[]int32           `db:"n_int32_arrays" bun:"n_int32_arrays,array"`
}

type TblDemoSetter struct {
	bun.BaseModel `bun:"table:tbl_demo,alias:u2"`
	ID            *int64                       `db:"-" bun:"-"`
	NStr          *sql.Null[string]            `db:"n_str" bun:"n_str"`
	VStr          *string                      `db:"v_str" bun:"v_str"`
	NEnum         *sql.Null[EnumTblBasicNEnum] `db:"n_enum" bun:"n_enum"`
	VEnum         *EnumTblBasicVEnum           `db:"v_enum" bun:"v_enum"`
	NInt32        *sql.Null[int32]             `db:"n_int32" bun:"n_int32"`
	VInt32        *int32                       `db:"v_int32" bun:"v_int32"`
	NTime         *sql.Null[time.Time]         `db:"n_time" bun:"n_time"`
	VTime         *time.Time                   `db:"v_time" bun:"v_time"`
	NUUID         *sql.Null[string]            `db:"n_uuid" bun:"n_uuid"`
	VUUID         *string                      `db:"v_uuid" bun:"v_uuid"`
	NJson         *sql.Null[map[string]any]    `db:"n_json" bun:"n_json"`
	VJson         *map[string]any              `db:"v_json" bun:"v_json"`
	VStrArrays    *pq.StringArray              `db:"v_str_arrays" bun:"v_str_arrays"`
	NStrArrays    *sql.Null[[]string]          `db:"n_str_arrays" bun:"n_str_arrays"`
	NBytes        *sql.Null[[]byte]            `db:"n_bytes" bun:"n_bytes"`
	VBytes        *pq.ByteaArray               `db:"v_bytes" bun:"v_bytes"`
}

type TblSimple struct {
	bun.BaseModel `bun:"table:tbl_simple,alias:u"`
	ID            int64           `bun:"id,pk"`
	Name          string          `bun:"name"`
	CreatedAt     time.Time       `bun:"created_at"`
	GUID          string          `bun:"guid"`
	DeletedAt     *time.Time      `bun:"deleted_at"`
	NJSON         *map[string]any `bun:"n_json" bun:"n_json"`
}

type TblSimpleSetter struct {
	bun.BaseModel `bun:"table:tbl_simple,alias:u"`
	ID            omit.Val[int64]              `bun:"id,pk,nullzero"`
	Name          omit.Val[string]             `bun:"name"`
	CreatedAt     omit.Val[time.Time]          `bun:"created_at"`
	GUID          omit.Val[string]             `bun:"guid"`
	DeletedAt     omitnull.Val[time.Time]      `bun:"deleted_at"`
	StrArrays     omitnull.Val[[]string]       `bun:"str_arrays"`
	NJSON         omitnull.Val[map[string]any] `bun:"n_json"`
}

type TblSimpleSetter2 struct {
	bun.BaseModel `bun:"table:tbl_simple,alias:u"`
	ID            *int64     `bun:"id,pk"`
	Name          *string    `bun:"name"`
	CreatedAt     *time.Time `bun:"created_at"`
	GUID          *string    `bun:"guid"`
	DeletedAt     *time.Time `bun:"deleted_at"`
	StrArrays     *[]string  `bun:"str_arrays"`
}

type TblSimpleColumns struct {
	ID        string
	Name      string
	CreatedAt string
	GUID      string
	DeletedAt string
}

var TblSimpleTable = schema.Table[TblSimple, TblSimpleColumns]{
	PrimaryKeys: []string{"id"},
	Model:       (*TblSimple)(nil),
	Name:        "tbl_simple",
	Alias:       "u",
	Columns: TblSimpleColumns{
		ID:        "id",
		Name:      "name",
		CreatedAt: "created_at",
		GUID:      "guid",
		DeletedAt: "deleted_at",
	},
}

func randomTblSimpleSetter(deletedAt *time.Time) *TblSimpleSetter {
	r := &TblSimpleSetter{
		Name:      omit.From(gofakeit.Name()),
		GUID:      omit.From(gofakeit.UUID()),
		CreatedAt: omit.From(time.Now()),
		DeletedAt: omitnull.FromPtr(deletedAt),
		//StrArrays: omitnull.From([]string{gofakeit.Name(), gofakeit.Animal(), gofakeit.City()}),
	}
	if v := gofakeit.IntRange(0, 6); int(v%2) == 0 {
		r.NJSON = omitnull.From(map[string]any{"name": gofakeit.Name(), "level": v})
	} else {

	}
	return r
}

func randomTblSimpleSetter2(deletedAt *time.Time) *TblSimpleSetter2 {
	r := &TblSimpleSetter2{
		Name:      ptr.Of(gofakeit.Name()),
		GUID:      ptr.Of(gofakeit.UUID()),
		CreatedAt: ptr.Of(time.Now()),
		DeletedAt: deletedAt,
		StrArrays: ptr.Of([]string{gofakeit.Name(), gofakeit.Animal(), gofakeit.City()}),
	}
	return r
}
