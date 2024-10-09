package sq

import (
	"context"
	"database/sql"
)

func RunInTx(ctx context.Context, db *sql.DB, opts *sql.TxOptions, fn func(context.Context, *sql.Tx) error) error {
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
