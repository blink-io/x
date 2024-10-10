package sq

import (
	"context"
	"database/sql"

	"github.com/bokwoon95/sq"
)

func RunInTx(ctx context.Context,
	db interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	}, opts *sql.TxOptions, fn func(context.Context, sq.DB) error) error {
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
