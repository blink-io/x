package sq

import "github.com/bokwoon95/sq"

type InsertMapper func() sq.InsertQuery
type UpdateMapper func() sq.UpdateQuery
type DeleteMapper func() sq.DeleteQuery
type QueryMapper[T any] func() (sq.SelectQuery, func(*sq.Row) T)

type Quote string

type mysql struct {
}

type postgres struct {
}
