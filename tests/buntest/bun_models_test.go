package buntest

import (
	"database/sql"
	"time"

	"github.com/blink-io/hyperbun/schema"
	"github.com/uptrace/bun"
)

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

var tblSimpleColumns = TblSimpleColumns{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	GUID:      "guid",
	DeletedAt: "deleted_at",
}

var TblSimpleTable = schema.Table[TblSimple, TblSimpleColumns]{
	PrimaryKeys: []string{"id"},
	Model:       (*TblSimple)(nil),
	Name:        "tbl_simple",
	Alias:       "u",
	Columns:     tblSimpleColumns,
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
