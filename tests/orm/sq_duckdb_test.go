package orm

import (
	"database/sql"
	"log/slog"
	"testing"

	"github.com/blink-io/x/ptr"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/marcboeker/go-duckdb"
	"github.com/stretchr/testify/require"
)

func getDuckDB() *sql.DB {
	sqliteOnce.Do(func() {
		setupDuckDBDialect()
	})
	dsn := "./orm-demo-duck.db?access_mode=READ_WRITE"
	db, err := sql.Open("duckdb", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func getDuckDBForSQ() *sql.DB {
	return getDuckDB()
}

func setupDuckDBDialect() {
	dialect := "duckdb"
	sq.DefaultDialect.Store(&dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}

func TestSq_DuckDB_User_Insert_ColumnMapper_1(t *testing.T) {
	db := getDuckDBForSQ()
	tbl := UserTable

	records := []User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	_, err := sq.Exec(db, sq.
		InsertInto(tbl).ColumnValues(func(col *sq.Column) {
		for _, r := range records {
			rptr := ptr.Of(r)
			rptr.ID = int(gofakeit.Int32())
			userInsertColumnMapper(col, *rptr)
		}
	}))
	require.NoError(t, err)
}

func TestSq_DuckDB_UserDevice_Insert_ColumnMapper_1(t *testing.T) {
	db := getDuckDBForSQ()
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

func TestSq_DuckDB_User_FetchAll_1(t *testing.T) {
	db := getDuckDBForSQ()
	tbl := UserTable

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, userModelRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSq_DuckDB_User_Delete_All(t *testing.T) {
	db := getDuckDBForSQ()
	tbl := UserTable

	_, err := sq.Exec(db, sq.
		DeleteFrom(tbl).
		Where(tbl.ID.GtInt(0)),
	)
	require.NoError(t, err)
}

func TestSq_DuckDB_Device_Insert_ColumnMapper_1(t *testing.T) {
	db := getDuckDBForSQ()
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

func TestSq_DuckDB_Device_FetchAll_1(t *testing.T) {
	db := getDuckDBForSQ()
	tbl := DeviceTable

	query := sq.From(tbl).Where(tbl.ID.GtInt(0)).Limit(100)
	records, err := sq.FetchAll(db, query, deviceRowMapper())

	require.NoError(t, err)
	require.NotNil(t, records)
}
