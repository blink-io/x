package orm

import (
	"database/sql"
	"log/slog"
	"sync"

	"github.com/blink-io/sq"
	"github.com/blink-io/sqx"
	"github.com/mattn/go-sqlite3"
	"github.com/qustavo/sqlhooks/v2"
	"github.com/qustavo/sqlhooks/v2/hooks/loghooks"
	"modernc.org/sqlite"
)

const (
	sqliteDriverName  = "sqlite-with-hooks"
	sqlite3DriverName = "sqlite3-with-hooks"
)

var sqliteOnce sync.Once

var sqliteDSN = "./orm_demo.db"

func init() {
	sql.Register(sqliteDriverName, sqlhooks.Wrap(&sqlite.Driver{}, loghooks.New()))
	sql.Register(sqlite3DriverName, sqlhooks.Wrap(&sqlite3.SQLiteDriver{}, loghooks.New()))
}

func MustGetSqliteDB() *sql.DB {
	db, err := sql.Open(sqliteDriverName, sqliteDSN)
	if err != nil {
		panic(err)
	}
	return db
}

func MustGetSqlite3DB() *sql.DB {
	db, err := sql.Open(sqlite3DriverName, sqliteDSN)
	if err != nil {
		panic(err)
	}
	return db
}

func getSqliteDB() *sql.DB {
	sqliteOnce.Do(func() {
		setupSqlite3Dialect()
	})

	return MustGetSqliteDB()
}

func getSqlite3DB() *sql.DB {
	sqliteOnce.Do(func() {
		setupSqlite3Dialect()
	})

	return MustGetSqlite3DB()
}

func setupSqlite3Dialect() {
	dialect := sq.DialectSQLite
	sqx.SetDefaultDialect(dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}

func GetSqliteDB() *sql.DB {
	return getSqliteDB()
}

func GetSqlite3DB() *sql.DB {
	return getSqlite3DB()
}
