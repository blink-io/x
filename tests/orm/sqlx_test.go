package orm

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestSqlx_1(t *testing.T) {
	db := sqlx.MustOpen("sqlite", "./orm_demo.db")

	m := new(Model)

	err := db.Get(m, "select 'heison' as name, sqlite_version() as version")
	require.NoError(t, err)

	fmt.Println(m)
}
