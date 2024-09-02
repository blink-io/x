package orm

import (
	"database/sql"
	"log/slog"
	"testing"

	"github.com/blink-io/x/tests/orm/nsqlite3"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
	//_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func init() {
	dialect := sq.DialectSQLite
	sq.DefaultDialect.Store(&dialect)
	slog.Info("Using dialect", "dialect", dialect)
}

func getSqlite3DB() *sql.DB {
	return nsqlite3.GetSQLiteDB("./orm_demo.db")
}

func TestSq_Sqlite3_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqlite3DB()
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

func TestSq_Sqlite3_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
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

func TestSq_Sqlite3_User_FetchAll_1(t *testing.T) {
	db := getSqliteDB()
	tbl := UserTableDef

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite3_User_Delete_All(t *testing.T) {
	db := getSqliteDB()
	tbl := UserTableDef

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.GtInt(0)),
	)
	require.NoError(t, err)
}
