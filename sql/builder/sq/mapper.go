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
)
