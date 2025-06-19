package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/opt/omitnull"
	_ "github.com/blink-io/opt/omitnull"
	"github.com/blink-io/sq"
	"github.com/stretchr/testify/require"
)

func selectDBNameAndVersion() sq.SelectQuery {
	sel := sq.Select(
		sq.DialectExpr("'no name' as name").
			DialectExpr(sq.DialectPostgres, "current_database() as name").
			DialectExpr(sq.DialectSQLite, "'sqlite has no name' as name"),

		sq.DialectExpr("'no version' as version").
			DialectExpr(sq.DialectPostgres, "version() as version").
			DialectExpr(sq.DialectSQLite, "sqlite_version() as version"),

		sq.DialectExpr("'no now datetime' as current").
			DialectExpr(sq.DialectPostgres, "to_char(now(), 'YYYY-MM-DD HH24:MI:SS') as current").
			DialectExpr(sq.DialectSQLite, " datetime('now','+8 hour') as current"),
	)
	return sel
}

func TestSq_Generic_DB_Select(t *testing.T) {
	sel := selectDBNameAndVersion()

	db1 := GetSqliteDB()
	m1, err1 := sq.FetchOne[Model](sq.Log(db1), sel.SetDialect(sq.DialectSQLite), func(r *sq.Row) Model {
		m := Model{
			Dialect: sq.DialectSQLite,
			Name:    r.String("name"),
			Version: r.String("version"),
			Current: r.String("current"),
		}
		return m
	})
	require.NoError(t, err1)
	require.NotNil(t, m1)

	fmt.Println("-------------------------------------------------------------------------------")

	fmt.Println(m1)

	fmt.Println("-------------------------------------------------------------------------------")

	db2 := GetPgDB()
	m2, err2 := sq.FetchOne[*Model](sq.Log(db2), sel.SetDialect(sq.DialectPostgres), func(r *sq.Row) *Model {
		m := &Model{
			Dialect: sq.DialectPostgres,
			Name:    r.String("name"),
			Version: r.String("version"),
			Current: r.String("current"),
		}
		return m
	})
	require.NoError(t, err2)
	require.NotNil(t, m2)

	fmt.Println("-------------------------------------------------------------------------------")

	fmt.Println(m2)

	fmt.Println("-------------------------------------------------------------------------------")
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

func TestModel_2(t *testing.T) {
	var mm = Model{
		Name:    "abc",
		Version: "efg",
		Current: "hij",
	}

	fmt.Printf("mm ptr 1: %p\n", &mm)

	mm1 := mm.WithName("new-name")
	mm = mm1

	fmt.Printf("mm ptr 2: %p\n", &mm)

	mm2 := mm.WithVersion("new-version")
	mm = mm2

	fmt.Printf("mm ptr 3: %p\n", &mm)

	mm3 := mm.WithName("after-name").
		WithVersion("after-version")
	mm = mm3

	fmt.Printf("mm ptr 3: %p\n", &mm)

	fmt.Printf("mm 1: %p\n", &mm1)
	fmt.Printf("mm 2: %p\n", &mm2)
	fmt.Printf("mm 3: %p\n", &mm3)
}
