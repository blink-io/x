package orm

import (
	"fmt"
	"log/slog"
	"testing"

	dbrslog "github.com/blink-io/x/orm/dbr/logger/slog"
	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestDBR_1(t *testing.T) {
	sl := dbrslog.New(slog.Default(), slog.LevelInfo)

	cc, err := dbr.Open("sqlite", "./orm_demo.db", sl)
	require.NoError(t, err)

	sess := cc.NewSession(nil)

	m := new(Model)

	err = sess.SelectBySql("select 'heison' as name, sqlite_version() as version").
		LoadOne(m)
	require.NoError(t, err)

	fmt.Println(m)
}
