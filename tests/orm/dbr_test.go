package orm

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/blink-io/x/orm/dbr/dialect"
	dbrslog "github.com/blink-io/x/orm/dbr/logger/slog"
	"github.com/gocraft/dbr/v2"
	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/require"
)

func getDBRConn() *dbr.Connection {
	sl := dbrslog.New(slog.Default(), slog.LevelInfo)

	db := getSqliteDBForSQ()
	dd := dialect.SQLite3
	cc := &dbr.Connection{
		DB:            db,
		EventReceiver: sl,
		Dialect:       dd,
	}

	return cc
}
func TestDBR_1(t *testing.T) {
	cc := getDBRConn()
	sess := cc.NewSession(nil)

	m := new(Model)

	var _ null.Time

	err := sess.SelectBySql("select datetime('now') as now, 'heison' as name, sqlite_version() as version").
		LoadOne(m)
	require.NoError(t, err)

	fmt.Println(m)
}

func TestDBR_User_Insert_1(t *testing.T) {
	cc := getDBRConn()

	sess := cc.NewSession(nil)

	m1 := randomUser()
	m2 := randomUser()

	_, err := sess.InsertInto("users").
		Columns([]string{"username", "guid", "score", "created_at", "updated_at"}...).
		Record(m1).
		Record(m2).
		Exec()
	require.NoError(t, err)

	fmt.Println("done")
}

func TestTime_1(t *testing.T) {
	tt := time.Now()
	bb, err := tt.MarshalText()
	kk := tt.String()
	require.NoError(t, err)
	fmt.Println(string(bb))
	fmt.Println(kk)
}
