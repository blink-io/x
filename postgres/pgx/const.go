package pgx

// Constants defined in pgxpool for runtime parameters.
const (
	PoolMinConns = "pool_min_conns"

	PoolMaxConns = "pool_max_conns"

	PoolMaxConnLifetime = "pool_max_conn_lifetime"

	PoolMaxConnLifetimeJitter = "pool_max_conn_lifetime_jitter"

	PoolMaxConnIdleTime = "pool_max_conn_idle_time"

	PoolHealthCheckPeriod = "pool_health_check_period"
)

// Constants defined in pgx for runtime parameters.
const (
	Timezone = "timezone"

	StatementCacheCapacity = "statement_cache_capacity"

	DescriptionCacheCapacity = "description_cache_capacity"

	DefaultQueryExecMode = "default_query_exec_mode"
)
