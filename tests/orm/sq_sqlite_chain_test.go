package orm

import (
	"context"
	"fmt"
	"testing"

	sqx "github.com/blink-io/x/sql/builder/sq"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Sqlite_Chain_1(t *testing.T) {
	var db sq.DB = getSqliteDBForSQ()
	f1 := func(db sq.DB) sq.DB {
		fmt.Println("Invoke f1")
		return db
	}
	f2 := func(db sq.DB) sq.DB {
		fmt.Println("Invoke f2")
		return db
	}
	f3 := func(db sq.DB) sq.DB {
		fmt.Println("Invoke f3")
		return db
	}
	chain := sqx.NewChain(f1, f2, f3)

	q := sq.Queryf("select sqlite_version() as ver")

	ver, err := sq.FetchOne(sq.Log(chain.Then(db)), q, func(r *sq.Row) string {
		return r.String("ver")
	})
	require.NoError(t, err)

	fmt.Println("sqlite version: ", ver)
}

func TestSq_Sqlite_Chain_2(t *testing.T) {
	rdb := getSqliteDBForSQ()
	db := sqx.InTx(rdb)

	err := db.RunInTx(ctx, nil, func(ctx context.Context, db sq.DB) error {
		q := sq.Queryf("select sqlite_version() as ver")

		ver, err := sq.FetchOne(sq.Log(db), q, func(r *sq.Row) string {
			return r.String("ver")
		})

		if err != nil {
			return nil
		}

		fmt.Println("sqlite version: ", ver)

		return nil
	})

	require.NoError(t, err)
}
