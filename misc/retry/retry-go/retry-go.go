package retry_go

import "github.com/avast/retry-go/v5"

var (
	BackOffDelay = retry.BackOffDelay
	New          = retry.New
)

// Do runs retryableFunc with the given options. It preserves the package-level
// convenience API of retry-go v4, which v5 replaced with retry.New(...).Do(...).
func Do(retryableFunc retry.RetryableFunc, opts ...retry.Option) error {
	return retry.New(opts...).Do(retryableFunc)
}

type (
	Retrier       = retry.Retrier
	Option        = retry.Option
	DelayContext  = retry.DelayContext
	DelayTypeFunc = retry.DelayTypeFunc
	OnRetryFunc   = retry.OnRetryFunc
	RetryIfFunc   = retry.RetryIfFunc
	RetryableFunc = retry.RetryableFunc
	Timer         = retry.Timer
)
