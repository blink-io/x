package goretry

import (
	"github.com/sethvargo/go-retry"
)

type (
	Backoff               = retry.Backoff
	BackoffFunc           = retry.BackoffFunc
	RetryFunc             = retry.RetryFunc
	RetryFuncValue[T any] = retry.RetryFuncValue[T]
)
