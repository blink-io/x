package buntest

import (
	"database/sql"
	"time"

	"github.com/blink-io/hyperbun/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
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
	ID        model.Column
	Name      model.Column
	CreatedAt model.Column
	GUID      model.Column
	DeletedAt model.Column
}

var tblSimpleColumns = TblSimpleColumns{
	ID:        model.Column("id"),
	Name:      model.Column("name"),
	CreatedAt: model.Column("created_at"),
	GUID:      model.Column("guid"),
	DeletedAt: model.Column("deleted_at"),
}

var TblSimpleTable = struct {
	Name    schema.Name
	Alias   schema.Name
	Model   *TblSimple
	Columns TblSimpleColumns
}{
	Name:    schema.Name("tbl_simple"),
	Alias:   schema.Name("ts1"),
	Model:   (*TblSimple)(nil),
	Columns: tblSimpleColumns,
}

func (s TblSimpleSetter) ColumnMapper(q *bun.InsertQuery) {
	if s.ID != nil {
		q.ColumnExpr(TblSimpleTable.Columns.ID.String(), *s.ID)
	}
	if s.Name != nil {
		q.ColumnExpr(TblSimpleTable.Columns.Name.String(), *s.Name)
	}
	if s.CreatedAt != nil {
		q.ColumnExpr(TblSimpleTable.Columns.CreatedAt.String(), *s.Name)
	}
	if s.GUID != nil {
		q.ColumnExpr(TblSimpleTable.Columns.GUID.String(), *s.Name)
	}
	if s.DeletedAt != nil && s.DeletedAt.Valid {
		q.ColumnExpr(TblSimpleTable.Columns.DeletedAt.String(), *s.Name)
	}
}
