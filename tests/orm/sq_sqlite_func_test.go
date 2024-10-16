package orm

import (
	"fmt"
	"testing"

	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Sqlite_Datetime_1(t *testing.T) {
	db := GetSqliteDB()

	defer db.Close()

	ss, err := sq.FetchOne(db, sq.Select(sq.Expr("{} || ' ::: ' || {} as now", sq.Expr("datetime()"), sq.Expr("unixepoch()"))), func(r *sq.Row) string {
		ss := r.String("now")
		return ss
	})
	require.NoError(t, err)

	fmt.Println(ss)
}
