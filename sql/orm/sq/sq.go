package sq

import (
	"log/slog"

	"github.com/bokwoon95/sq"
)

type (
	DB        = sq.DB
	Predicate = sq.Predicate
	Query     = sq.Query
	Row       = sq.Row
	SQLWriter = sq.SQLWriter

	JSONMap map[string]any
)

func SetDefaultDialect(dialect string) {
	switch dialect {
	case sq.DialectPostgres:
	case sq.DialectSQLite:
	case sq.DialectSQLServer:
	case sq.DialectMySQL:
		sq.DefaultDialect.Store(&dialect)
	default:
		slog.Warn("")
	}
}

func UnsetDefaultDialect() {
	sq.DefaultDialect.Store(nil)
}

func Quote(dialect string) string {
	return ""
}
