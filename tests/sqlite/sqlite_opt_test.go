package sqlite

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/dbk"
	"github.com/blink-io/x/sql/dbm"
	"github.com/blink-io/x/sql/dbq"
	"github.com/blink-io/x/sql/dbr"
	"github.com/blink-io/x/sql/dbs"
	"github.com/blink-io/x/sql/dbx"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/stephenafamo/bob"
)

func init() {
	sqliteDialectOpts := sqlite3.DialectOptions()
	sqliteDialectOpts.SupportsReturn = true
	goqu.RegisterDialect(xsql.DialectSQLite, sqliteDialectOpts)
}

func getSqliteSqlDB() *sql.DB {
	db, err := xsql.NewSqlDB(sqliteCfg())
	//db.AddQueryHook(logging.Func(log.Printf))
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDB() *xdb.DB {
	db, err := xdb.New(sqliteCfg(), dbOpts()...)

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBQ() *dbq.DB {
	db, err := dbq.New(sqliteCfg())
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBR() *dbr.DB {
	db, err := dbr.New(sqliteCfg(),
		dbr.WithEventReceiver(dbr.NewTimingEventReceiver()),
	)

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBM() *dbm.DB {
	db, err := dbm.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBS() *dbs.DB {
	db, err := dbs.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBK() *dbk.DB {
	db, err := dbk.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBX() *dbx.DB {
	ops := []dbx.Option{
		dbx.ExecWrappers(
			bob.Debug,
			func(exec bob.Executor) bob.Executor {
				return dbx.ExecOnError(exec, func(e error) error {
					return xsql.WrapError(e)
				})
			}),
	}
	db, err := dbx.New(sqliteCfg(), ops...)
	if err != nil {
		panic(err)
	}

	return db
}
