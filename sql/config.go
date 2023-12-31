package sql

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/blink-io/x/sql/hooks"
	"go.opentelemetry.io/otel/attribute"

	"github.com/uptrace/bun"
)

type QueryHook = bun.QueryHook

type Config struct {
	Context         context.Context
	Dialect         string
	Network         string
	Host            string
	Port            int
	Name            string
	User            string
	Password        string
	TLSConfig       *tls.Config
	Params          map[string]string
	DialTimeout     time.Duration
	ConnInitSQL     string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ValidationSQL   string
	DriverHooks     []hooks.Hooks

	Loc        *time.Location
	Debug      bool
	Collation  string
	ClientName string
	WithOTel   bool
	Attrs      []attribute.KeyValue
	Logger     func(format string, args ...any)
	dsn        string
	accessor   string
}

func setupConfig(c *Config) *Config {
	if c == nil {
		c = new(Config)
	}
	c.SetDefaults()
	return c
}

func (c *Config) SetDefaults() {
	if c == nil {
		return
	}

	if c.Context == nil {
		c.Context = context.Background()
	}
	if len(c.Network) == 0 {
		c.Network = "tcp"
	}
	if c.Loc == nil {
		c.Loc = time.Local
	}
}

func (c *Config) Validate(ctx context.Context) error {
	return nil
}

func newDBInfo(c *Config) DBInfo {
	return DBInfo{
		Name:    c.Name,
		Dialect: c.Dialect,
	}
}
