package buntest

import (
	"database/sql"
	"github.com/alexlast/bunzap"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunslog"
	"go.uber.org/zap"
	"log"
	"log/slog"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/blink-io/hypersql"
	pgparams "github.com/blink-io/hypersql/postgres/params"
	"github.com/blink-io/sq"
	"github.com/blink-io/sqx"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2/hooks/loghooks"
)

const (
	pgDriverName = "pgx-with-hooks"
)

var pgOnce sync.Once

func GetPgDB() *sql.DB {
	return getPgDB()
}

func getPgDBForSqlx() *sqlx.DB {
	return sqlx.NewDb(getPgDB(), pgDriverName)
}

func getPgDB() *sql.DB {
	pgOnce.Do(func() {
		setupPgDialect()
	})

	host := "192.168.50.88"
	port := 5432
	//sql.Register(pgDriverName, sqlhooks.Wrap(stdlib.GetDefaultDriver(), loghooks.New()))
	if strings.EqualFold(runtime.GOOS, "darwin") {
		host = "localhost"
		port = 15432
	}

	c := &hypersql.Config{
		Dialect:  hypersql.DialectPostgres,
		Host:     host,
		Port:     port,
		User:     "test",
		Password: "test",
		Name:     "test",
		Params: hypersql.ConfigParams{
			pgparams.ConnParams.ApplicationName: "go-client-test-n1",
			"TimeZone":                          "Asia/Shanghai",
		},
		DriverHooks: hypersql.DriverHooks{
			loghooks.New(),
		},
		Loc: time.Local,
	}

	//dsn := "postgres://test:test@:5432/test?sslmode=disable&TimeZone=Asia/Shanghai"
	//db, err := sql.Open(pgDriverName, dsn)
	db, err := hypersql.NewSqlDB(c)
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

func setupBunHooks(bundb *bun.DB) {
	h1 := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	h2 := bunslog.NewQueryHook(
		bunslog.WithQueryLogLevel(slog.LevelInfo),
		bunslog.WithSlowQueryLogLevel(slog.LevelWarn),
		bunslog.WithLogFormat(func(event *bun.QueryEvent) []slog.Attr {
			return []slog.Attr{
				slog.String("operation", event.Operation()),
			}
		}),
	)
	h3 := bunzap.NewQueryHook(bunzap.QueryHookOptions{
		SlowDuration: 5 * time.Second,
		Logger:       zap.L(),
	})
	bundb.AddQueryHook(h1)
	bundb.AddQueryHook(h2)
	bundb.AddQueryHook(h3)
}
