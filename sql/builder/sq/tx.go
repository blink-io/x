package sq

import (
	"context"
	"database/sql"

	"github.com/bokwoon95/sq"
)

type (
	Txer interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	}

	RunInTxer interface {
		RunInTx(context.Context, *sql.TxOptions, func(context.Context, sq.DB) error) error
	}

	runInTxDB struct {
		Txer
		sq.DB
	}
)

func (db runInTxDB) RunInTx(ctx context.Context, opts *sql.TxOptions, fn func(context.Context, sq.DB) error) error {
	return RunInTx(ctx, db, opts, fn)
}

func InTx(db interface {
	sq.DB
	Txer
}) interface {
	sq.DB
	Txer
	RunInTxer
} {
	return runInTxDB{DB: db, Txer: db}
}

func RunInTx(ctx context.Context,
	db Txer, opts *sql.TxOptions, fn func(context.Context, sq.DB) error) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	done = true
	return tx.Commit()
}
