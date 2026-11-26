package goretry

import (
	"github.com/sethvargo/go-retry"
)

type (
	// Backoff returns the wait before the next retry, or stop=true to end
	// retrying. go-retry v0.7 redefined it from an interface to this func
	// type, which now also receives the number of retries made so far.
	Backoff = retry.Backoff

	RetryFunc             = retry.RetryFunc
	RetryFuncValue[T any] = retry.RetryFuncValue[T]
)
