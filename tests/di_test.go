package tests

import (
	"errors"
	"fmt"
	"testing"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/samber/do/v2"
	"github.com/stretchr/testify/require"
)

var yesErr = errors.New("yes error")

func TestDo_1(t *testing.T) {
	// create DI container
	i := do.New()

	// inject both services into DI container
	do.Provide[*xsql.DB](i, NewDB)
	do.Provide[*xsql.DBP](i, NewDBPWithErr)
	do.Provide[*xsql.Config](i, NewOptions)

	uname := "uni-opts"
	do.ProvideNamedTransient(i, uname, NewOptions)

	opt1 := do.MustInvoke[*xsql.Config](i)
	opt2 := do.MustInvoke[*xsql.Config](i)
	require.NotNil(t, opt1)
	require.NotNil(t, opt2)

	uopt1 := do.MustInvokeNamed[*xsql.Config](i, uname)
	uopt2 := do.MustInvokeNamed[*xsql.Config](i, uname)
	require.NotNil(t, uopt1)
	require.NotNil(t, uopt2)

	db, err := do.Invoke[*xsql.DB](i)
	require.NoError(t, err)

	dbx, err2 := do.Invoke[*xsql.DBX](i)
	require.Nil(t, dbx)
	require.Error(t, err2)

	dbp, err3 := do.Invoke[*xsql.DBP](i)
	require.Nil(t, dbp)
	require.Equal(t, yesErr, err3)

	row := db.QueryRow("select sqlite_version()")
	var str string
	require.NoError(t, row.Scan(&str))

	fmt.Println("Result: ", str)

	i.Shutdown()
}

func NewDBPWithErr(i do.Injector) (*xsql.DBP, error) {
	return nil, yesErr
}

func NewDB(i do.Injector) (*xsql.DB, error) {
	return xsql.NewDB(do.MustInvoke[*xsql.Config](i))
}

func NewOptions(i do.Injector) (*xsql.Config, error) {
	var opt = &xsql.Config{
		Dialect: xsql.DialectSQLite,
		Host:    sqlitePath,
		DOptions: []xsql.DialectOption{
			xsql.DialectWithLoc(time.Local),
		},
		Loc: time.Local,
	}
	return opt, nil
}
