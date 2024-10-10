package sq

import (
	"context"

	"github.com/bokwoon95/sq"
)

type (
	Mapper[T sq.Table, M any, S any] interface {
		Table() T

		InsertT(context.Context, ...S) func(*sq.Column)

		UpdateT(context.Context, S) func(*sq.Column)

		QueryT(context.Context) func(*sq.Row) M
	}

	Executor[M any, S any] interface {
		Insert(ctx context.Context, db sq.DB, ss ...S) (sq.Result, error)

		Update(ctx context.Context, db sq.DB, where sq.Predicate, s S) (sq.Result, error)

		Delete(ctx context.Context, db sq.DB, where sq.Predicate) (sq.Result, error)

		One(ctx context.Context, db sq.DB, where sq.Predicate) (M, error)

		All(ctx context.Context, db sq.DB, where sq.Predicate) ([]M, error)
	}
)
