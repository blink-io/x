package multierr

import (
	"go.uber.org/multierr"
)

var (
	Append = multierr.Append

	AppendInto = multierr.AppendInto

	AppendFunc = multierr.AppendFunc

	Close = multierr.Close

	Combine = multierr.Combine

	Errors = multierr.Errors

	Every = multierr.Every
)

type (
	Invoke = multierr.Invoke

	Invoker = multierr.Invoker
)
