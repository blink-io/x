package buntest

import (
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestPg_Users_SelectJoin_1(t *testing.T) {
	db := getPgDB()
	cols := []string{
		TblSimpleTable.Columns.GUID,
		TblSimpleTable.Columns.Name,
		TblSimpleTable.Columns.CreatedAt,
	}
	vals := []any{
		gofakeit.UUID(),
		gofakeit.Name(),
		time.Now(),
	}
	_, err := sq.Insert(TblSimpleTable.Name).
		Columns(cols...).
		Values(vals...).
		RunWith(db).
		PlaceholderFormat(sq.Dollar).
		Exec()
	require.Nil(t, err)
}
