package orm

import (
	"database/sql"
	"testing"

	"github.com/blink-io/x/tests/orm/msqlite3"
	"github.com/bokwoon95/sq"
	"github.com/stretchr/testify/require"
)

func getSqliteDB() *sql.DB {
	sqliteOnce.Do(func() {
		setupSqlite3Dialect()
	})

	return msqlite3.GetSqliteDB(sqliteDSN)
}

func getSqliteDBForSQ() *sql.DB {
	return getSqliteDB()
}

func TestSq_Sqlite_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserTable

	records := []User{
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

func TestSq_Sqlite_User_FetchAll_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserTable

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_Sqlite_User_Delete_All(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.GtInt(0)),
	)
	require.NoError(t, err)
}

func TestSq_Sqlite_Device_Insert_ColumnMapper_1(t *testing.T) {
	db := getSqliteDBForSQ()
	tbl := DeviceTable

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
	db := getSqliteDBForSQ()
	tbl := DeviceTable

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, deviceRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}
