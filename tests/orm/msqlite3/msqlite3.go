package msqlite3

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetSqliteDB(dsn string) *sql.DB {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
