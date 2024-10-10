package orm

import (
	"database/sql"
	"log/slog"
	"sync"

	"github.com/blink-io/sqx"
	"github.com/bokwoon95/sq"
	"github.com/qustavo/sqlhooks/v2"
	"github.com/qustavo/sqlhooks/v2/hooks/loghooks"
	"modernc.org/sqlite"
)

const (
	sqliteDriverName = "sqlite-with-hooks"
)

var sqliteOnce sync.Once

var sqliteDSN = "./orm_demo.db"

func GetSqliteDB(dsn string) *sql.DB {
	sql.Register(sqliteDriverName, sqlhooks.Wrap(&sqlite.Driver{}, loghooks.New()))

	db, err := sql.Open(sqliteDriverName, dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func getSqliteDB() *sql.DB {
	sqliteOnce.Do(func() {
		setupSqlite3Dialect()
	})

	return GetSqliteDB(sqliteDSN)
}

func setupSqlite3Dialect() {
	dialect := sq.DialectSQLite
	sqx.SetDefaultDialect(dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}

func getSqliteDBForSQ() *sql.DB {
	return getSqliteDB()
}
