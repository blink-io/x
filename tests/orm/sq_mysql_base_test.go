package orm

import (
	"database/sql"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/blink-io/hypersql"
	"github.com/blink-io/sq"
	"github.com/blink-io/sqx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qustavo/sqlhooks/v2/hooks/loghooks"
)

const (
	mysqlDriverName = "mysql-with-hooks"
)

var mysqlOnce sync.Once

func getMysqlDBForSQ() *sql.DB {
	return getMysqlDB()
}

func getMysqlDB() *sql.DB {
	mysqlOnce.Do(func() {
		setupMysqlDialect()
	})

	//sql.Register(mysqlDriverName, sqlhooks.Wrap(&mysql.MySQLDriver{}, loghooks.New()))

	c := &hypersql.Config{
		Dialect:  hypersql.DialectMySQL,
		Host:     "192.168.50.88",
		Port:     3306,
		User:     "test",
		Password: "test",
		Name:     "test",
		Params: hypersql.ConfigParams{
			"ParseTime": "true",
			"TimeZone":  "Asia/Shanghai",
		},
		DriverHooks: hypersql.DriverHooks{
			loghooks.New(),
		},
		Loc: time.Local,
	}

	db, err := hypersql.NewSqlDB(c)

	//dsn := "test:test@tcp(192.168.50.88:3306)/test?parseTime=true&loc=Local"
	//db, err := sql.Open(mysqlDriverName, dsn)
	if err != nil {
		log.Fatalf("failed to open mysql db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping mysql db: %v", err)
	}

	return db
}

func setupMysqlDialect() {
	dialect := sq.DialectMySQL
	sqx.SetDefaultDialect(dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}
