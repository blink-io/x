package orm

import (
	"database/sql"
	"log"
	"log/slog"
	"sync"

	"github.com/bokwoon95/sq"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlOnce sync.Once

func getMysqlDBForSQ() *sql.DB {
	return getMysqlDB()
}

func getMysqlDB() *sql.DB {
	mysqlOnce.Do(func() {
		setupMysqlDialect()
	})

	dsn := "test:test@tcp(192.168.50.88:3306)/test?parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
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
	sq.DefaultDialect.Store(&dialect)
	slog.Info("Setup database dialect", "dialect", dialect)
}
