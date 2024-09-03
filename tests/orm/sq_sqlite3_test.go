package orm

import (
	"database/sql"
	"log/slog"
	"sync"
	"testing"

	"github.com/blink-io/x/tests/orm/nsqlite3"

	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

var sqliteOnce sync.Once

var sqliteDSN = "./orm_demo.db"

func getSqlite3DB() *sql.DB {
	sqliteOnce.Do(func() {
		setupSqlite3Dialect()
	})

	return nsqlite3.GetSQLiteDB(sqliteDSN)
}

func getSqlite3DBForSQ() sq.DB {
	return sq.Log(getSqlite3DB())
}

func setupSqlite3Dialect() {
	dialect := sq.DialectSQLite
	sq.DefaultDialect.Store(&dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}

func TestSq_Sqlite3_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqlite3DBForSQ()
	tbl := UserTable

	records := []*User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			userInsertColumnMapper(col, r)
		}
	}))
	require.NoError(t, err)
}

func TestSq_Sqlite3_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserDeviceTable

	records := []*UserDevice{
		randomUserDevice(),
		randomUserDevice(),
		randomUserDevice(),
	}

	rt, err := sq.Exec(sq.Log(db), sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			userDeviceInsertColumnMapper(col, r)
		}
	}))
	require.NoError(t, err)
	require.NotNil(t, rt)
}

func TestSq_Sqlite3_User_FetchAll_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserTable

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite3_User_Delete_All(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.GtInt(0)),
	)
	require.NoError(t, err)
}
