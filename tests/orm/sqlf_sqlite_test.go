package orm

import (
	"database/sql"
	"testing"

	"github.com/leporo/sqlf"
	"github.com/stretchr/testify/require"
)

func TestSqlf_Sqlite_Select_Ver_1(t *testing.T) {
	db := getSqliteDB()

	m := new(Model)

	err := sqlf.
		Select("'heison' as name").To(&m.Name).
		Select("sqlite_version() as version").To(&m.Version).
		QueryAndClose(ctx, db, func(rows *sql.Rows) {
		})
	require.NoError(t, err)
}
