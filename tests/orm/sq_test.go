package orm

import (
	"database/sql"
	"fmt"
	"testing"

	qm "github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

type MODEL struct {
	qm.TableStruct
	Version qm.String
}

type Model struct {
	Name    string `db:"name"`
	Version string `db:"version"`
}

func (m Model) String() string {
	return "name=" + m.Name + " version=" + m.Version
}

var _ fmt.Stringer = (*Model)(nil)

func TestSq_1(t *testing.T) {
	db, err := sql.Open("sqlite", "./orm_demo.db")
	require.NoError(t, err)

	qr := qm.SQLite.Queryf("select 'heison' as name, sqlite_version() as version")
	m, err := qm.FetchOne[*Model](db, qr, func(row *qm.Row) *Model {
		return &Model{
			Name:    row.String("name"),
			Version: row.String("version"),
		}
	})
	require.NoError(t, err)
	require.NotNil(t, m)

	fmt.Println(m)
}
