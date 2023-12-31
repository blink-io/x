//go:build !nosqlite

package sql

import (
	"context"

	"github.com/blink-io/x/bun/dialect/sqlitedialect"

	"github.com/glebarez/go-sqlite"
	"github.com/uptrace/bun/schema"
)

const (
	DialectSQLite = "sqlite"
)

func init() {
	dn := DialectSQLite
	drivers[dn] = &sqlite.Driver{}
	dialectors[dn] = NewSQLiteDialect
	dsners[dn] = SQLiteDSN
}

func NewSQLiteDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	dopt := applyDialectOptions(ops...)
	sops := make([]sqlitedialect.Option, 0)
	if dopt.loc != nil {
		sops = append(sops, sqlitedialect.Location(dopt.loc))
	}
	return sqlitedialect.New(sops...)
}

func SQLiteDSN(ctx context.Context, c *Config) (string, error) {
	dsn := c.Host
	return dsn, nil
}
