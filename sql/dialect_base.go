package sql

import (
	"context"
	"database/sql/driver"
	"strings"

	"github.com/blink-io/x/sql/driver/hooks"
	"github.com/life4/genesis/slices"
)

type (
	Dsner = func(context.Context, *Config) (string, error)

	ConnectorFunc func(ctx context.Context, c *Config) (driver.Connector, error)

	//GetDriverFunc func(dialect string) (driver.Driver, error)

	//GetDSNFunc func(dialect string) (Dsner, error)
)

var (
	//drivers = make(map[string]GetDriverFunc)

	//dsners = make(map[string]GetDSNFunc)

	connectors = make(map[string]ConnectorFunc)
)

func GetFormalDialect(dialect string) string {
	if d, ok := IsCompatibleDialect(dialect); ok {
		return d
	}
	return ""
}

// IsCompatibleDialect checks
func IsCompatibleDialect(dialect string) (string, bool) {
	if IsCompatiblePostgresDialect(dialect) {
		return DialectPostgres, true
	} else if IsCompatibleMySQLDialect(dialect) {
		return DialectMySQL, true
	} else if IsCompatibleSQLiteDialect(dialect) {
		return DialectSQLite, true
	}
	return "", false
}

func isCompatibleDialectIn(dialect string, compatibleDialects []string) bool {
	dialect = strings.ToLower(dialect)
	i := slices.FindIndex(compatibleDialects, func(i string) bool {
		return i == dialect
	})
	return i > -1
}

func wrapDriverHooks(drv driver.Driver, drvHooks ...hooks.Hooks) driver.Driver {
	if len(drvHooks) > 0 {
		drv = hooks.Wrap(drv, hooks.Compose(drvHooks...))
	}
	return drv
}
