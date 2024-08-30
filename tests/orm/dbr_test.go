package orm

import (
	"fmt"
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestDBR_1(t *testing.T) {
	cc, err := dbr.Open("sqlite", "./orm_demo.db", nil)
	require.NoError(t, err)

	m := new(Model)

	r := cc.QueryRow("select 'heison' as name, sqlite_version() as version")
	require.NoError(t, r.Scan(&m.Name, &m.Version))

	fmt.Println(m)
}
