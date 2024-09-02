package orm

import (
	"database/sql"
	"log/slog"
	"testing"

	"github.com/blink-io/x/tests/orm/msqlite3"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
	//_ "github.com/ncruces/go-sqlite3/driver"
	//_ "github.com/ncruces/go-sqlite3/embed"
)

func init() {
	dialect := sq.DialectSQLite
	sq.DefaultDialect.Store(&dialect)
	slog.Info("Using dialect", "dialect", dialect)
}

func getSqliteDB() *sql.DB {
	dsn := "./orm_demo.db"
	return msqlite3.GetSqliteDB(dsn)
}

func TestSq_Sqlite_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDB()
	tbl := UserTableDef

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

func TestSq_Sqlite_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDB()
	tbl := UserDeviceTableDef

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

func TestSq_Sqlite_User_FetchAll_1(t *testing.T) {
	db := getSqliteDB()
	tbl := UserTableDef

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite_User_Delete_All(t *testing.T) {
	db := getSqliteDB()
	tbl := UserTableDef

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.GtInt(0)),
	)
	require.NoError(t, err)
}

func TestSq_Sqlite_Device_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDB()
	tbl := DeviceTableDef

	records := []*Device{
		randomDevice(),
		randomDevice(),
		randomDevice(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			deviceInsertColumnMapper(col, r)
		}
	}))
	require.NoError(t, err)
}

func TestSq_Sqlite_Device_FetchAll_1(t *testing.T) {
	db := getSqliteDB()
	tbl := DeviceTableDef

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, deviceRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}
