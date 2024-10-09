package sqlx

import "github.com/jmoiron/sqlx"

type (
	Conn      = sqlx.Conn
	DB        = sqlx.DB
	Queryer   = sqlx.Queryer
	NamedStmt = sqlx.NamedStmt
)
