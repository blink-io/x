package puddle

import "github.com/jackc/puddle/v2"

var (
	ErrClosedPool = puddle.ErrClosedPool

	ErrNotAvailable = puddle.ErrNotAvailable
)

type (
	Stat = puddle.Stat

	Pool[T any]        = puddle.Pool[T]
	Constructor[T any] = puddle.Constructor[T]
	Config[T any]      = puddle.Config[T]
	Destructor[T any]  = puddle.Destructor[T]
	Resource[T any]    = puddle.Resource[T]
)
