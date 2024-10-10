package orm

import (
	"github.com/blink-io/hyperbun"
)

func getPgDBForBun() *hyperbun.DB {
	db, err := hyperbun.NewFromSqlDB(getPgDB(), hyperbun.DialectPostgres)
	if err != nil {
		panic(err)
	}
	return db
}
