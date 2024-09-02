package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/aarondl/opt/omitnull"
	_ "github.com/aarondl/opt/omitnull"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func TestSq_Generic_DB_Select(t *testing.T) {
	sel := sq.Select(
		sq.DialectExpr("'no name' as name").
			DialectExpr(sq.DialectPostgres, "current_database() as name").
			DialectExpr(sq.DialectSQLite, "'sqlite has no name' as name"),

		sq.DialectExpr("'no version' as version").
			DialectExpr(sq.DialectPostgres, "version() as version").
			DialectExpr(sq.DialectSQLite, "sqlite_version() as version"),

		sq.DialectExpr("'no now datetime' as current").
			DialectExpr(sq.DialectPostgres, "now() as current").
			DialectExpr(sq.DialectSQLite, " datetime('now','+8 hour') as current"),
	)

	db1 := getSqliteDB()
	m1, err1 := sq.FetchOne[Model](db1, sel.SetDialect(sq.DialectSQLite), func(r *sq.Row) Model {
		m := Model{
			Name:    r.String("name"),
			Version: r.String("version"),
			Current: r.String("current"),
		}
		return m
	})
	require.NoError(t, err1)
	require.NotNil(t, m1)

	fmt.Println(m1)

	db2 := getPgDB()
	m2, err2 := sq.FetchOne[*Model](db2, sel.SetDialect(sq.DialectPostgres), func(r *sq.Row) *Model {
		m := &Model{
			Name:    r.String("name"),
			Version: r.String("version"),
		}
		return m
	})
	require.NoError(t, err2)
	require.NotNil(t, m2)

	fmt.Println(m2)
}

func TestPtr_1(t *testing.T) {
	m := Model{
		Name:    "OK",
		Version: "v1.0.0",
	}

	fmt.Println("before: ", m)

	mptr := &m

	mptr.Name = "Fuck CCP"

	fmt.Println("after: ", m)

	m2 := m

	fmt.Println("m2: ", m2)
}

func TestPtr_2(t *testing.T) {
	m1 := Model{
		Name:    "Before",
		Version: "v1.0.0",
	}
	fmt.Println("m1: ", m1)

	m2 := m1
	fmt.Println("before m2: ", m2)

	m2ptr := &m2
	m2ptr.Name = "After"
	fmt.Println("after m2: ", m2)

	m2ptr2 := &m2
	m2ptr3 := m2ptr

	m2ptr3.Name = "After After"

	fmt.Println("after after m2: ", m2)

	fmt.Printf("m1 ptr: %p\n", &m1)
	fmt.Printf("m2 ptr: %p\n", &m2)
	fmt.Printf("m2 ptr2: %p\n", m2ptr2)
	fmt.Printf("m2 ptr3: %p\n", m2ptr3)
}

func TestEmitnull_1(t *testing.T) {
	tt := omitnull.From(time.Now())
	require.NotNil(t, tt)
}
