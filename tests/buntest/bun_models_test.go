package buntest

import (
	"database/sql"
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

type TblDemo struct {
	bun.BaseModel `bun:"table:tbl_demo,alias:u1"`
	ID            int64                       ` db:"-" bun:"-"`
	NStr          sql.Null[string]            ` db:"n_str" bun:"n_str"`
	VStr          string                      ` db:"v_str" bun:"v_str"`
	NEnum         sql.Null[EnumTblBasicNEnum] ` db:"n_enum" bun:"n_enum"`
	VEnum         EnumTblBasicVEnum           ` db:"v_enum" bun:"v_enum"`
	NInt32        sql.Null[int32]             ` db:"n_int32" bun:"n_int32"`
	VInt32        int32                       ` db:"v_int32" bun:"v_int32"`
	NTime         sql.Null[time.Time]         ` db:"n_time" bun:"n_time"`
	VTime         time.Time                   ` db:"v_time" bun:"v_time"`
	NUUID         sql.Null[string]            ` db:"n_uuid" bun:"n_uuid"`
	VUUID         string                      ` db:"v_uuid" bun:"v_uuid"`
	NJson         sql.Null[map[string]any]    ` db:"n_json" bun:"n_json"`
	VJson         map[string]any              ` db:"v_json" bun:"v_json"`
	VStrArrays    pq.StringArray              ` db:"v_str_arrays" bun:"v_str_arrays"`
	NStrArrays    sql.Null[[]string]          ` db:"n_str_arrays" bun:"n_str_arrays"`
	NBytes        sql.Null[[]byte]            ` db:"n_bytes" bun:"n_bytes"`
	VBytes        pq.ByteaArray               ` db:"v_bytes" bun:"v_bytes"`
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
	ID            int64               `bun:"id,pk"`
	Name          string              `bun:"name"`
	CreatedAt     time.Time           `bun:"created_at"`
	GUID          string              `bun:"guid"`
	DeletedAt     sql.Null[time.Time] `bun:"deleted_at"`
}
type TblSimpleSetter struct {
	bun.BaseModel `bun:"table:tbl_simple,alias:u"`
	ID            *int64               `bun:"id,pk"`
	Name          *string              `bun:"name"`
	CreatedAt     *time.Time           `bun:"created_at"`
	GUID          *string              `bun:"guid"`
	DeletedAt     *sql.Null[time.Time] `bun:"deleted_at"`
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

func (s TblSimpleSetter) ColumnMapper(q *bun.InsertQuery) {
	if s.ID != nil {
		q.ColumnExpr(TblSimpleTable.Columns.ID, *s.ID)
	}
	if s.Name != nil {
		q.ColumnExpr(TblSimpleTable.Columns.Name, *s.Name)
	}
	if s.CreatedAt != nil {
		q.ColumnExpr(TblSimpleTable.Columns.CreatedAt, *s.Name)
	}
	if s.GUID != nil {
		q.ColumnExpr(TblSimpleTable.Columns.GUID, *s.Name)
	}
	if s.DeletedAt != nil && s.DeletedAt.Valid {
		q.ColumnExpr(TblSimpleTable.Columns.DeletedAt, *s.Name)
	}
}
