package buntest

import (
	"database/sql"
	"log"
	"log/slog"
	"runtime"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunslog"

	"github.com/blink-io/hypersql"
	pgparams "github.com/blink-io/hypersql/postgres/params"
)

func GetPgDB() *sql.DB {
	return getPgDB()
}

func getPgDB() *sql.DB {
	host := "192.168.50.88"
	port := 5432
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
		Loc: time.Local,
	}

	db, err := hypersql.NewSqlDB(c)
	if err != nil {
		log.Fatalf("failed to open pg db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping pg db: %v", err)
	}

	return db
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
	h3 := bunslog.NewQueryHook(
		bunslog.WithQueryLogLevel(slog.LevelInfo),
		bunslog.WithSlowQueryLogLevel(slog.LevelWarn),
		bunslog.WithSlowQueryThreshold(5*time.Second),
	)
	bundb.WithQueryHook(h1)
	bundb.WithQueryHook(h2)
	bundb.WithQueryHook(h3)
}
