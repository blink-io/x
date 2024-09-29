package nsqlite3

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func GetSQLiteDB(dsn string) *sql.DB {
	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		panic(err)
	}
	return db
}
