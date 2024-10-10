package orm

import (
	"database/sql"
	"log"
	"log/slog"
	"sync"

	sqx "github.com/blink-io/x/sql/builder/sq"
	"github.com/bokwoon95/sq"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
	"github.com/qustavo/sqlhooks/v2/hooks/loghooks"
)

const (
	pgDriverName = "pgx-with-hooks"
)

var pgOnce sync.Once

func getPgDBForSQ() *sql.DB {
	return getPgDB()
}

func getPgDBForSqlx() *sqlx.DB {
	return sqlx.NewDb(getPgDB(), pgDriverName)
}

func getPgDB() *sql.DB {
	pgOnce.Do(func() {
		setupPgDialect()
	})

	sql.Register(pgDriverName, sqlhooks.Wrap(stdlib.GetDefaultDriver(), loghooks.New()))

	dsn := "postgres://test:test@192.168.50.88:5432/test?sslmode=disable&TimeZone=Asia/Shanghai"
	db, err := sql.Open(pgDriverName, dsn)
	if err != nil {
		log.Fatalf("failed to open pg db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping pg db: %v", err)
	}

	return db
}

func setupPgDialect() {
	dialect := sq.DialectPostgres
	sqx.SetDefaultDialect(dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}
