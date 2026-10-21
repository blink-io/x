package retry_go

import (
	"github.com/avast/retry-go/v5"
)

var (
	Attempts                      = retry.Attempts
	AttemptsForError              = retry.AttemptsForError
	Context                       = retry.Context
	Delay                         = retry.Delay
	DelayType                     = retry.DelayType
	LastErrorOnly                 = retry.LastErrorOnly
	MaxDelay                      = retry.MaxDelay
	MaxJitter                     = retry.MaxJitter
	OnRetry                       = retry.OnRetry
	RetryIf                       = retry.RetryIf
	UntilSucceeded                = retry.UntilSucceeded
	WithTimer                     = retry.WithTimer
	WrapContextErrorWithLastError = retry.WrapContextErrorWithLastError

	BackOffDelay = retry.BackOffDelay
	FixedDelay   = retry.FixedDelay
	RandomDelay  = retry.RandomDelay

	CombineDelay = retry.CombineDelay

	FullJitterBackoffDelay = retry.FullJitterBackoffDelay

	IsRecoverable = retry.IsRecoverable
	Unrecoverable = retry.Unrecoverable
)

type (
	Retrier                = retry.Retrier
	RetrierWithData[T any] = retry.RetrierWithData[T]

	Option = retry.Option

	DelayContext  = retry.DelayContext
	DelayTypeFunc = retry.DelayTypeFunc

	OnRetryFunc = retry.OnRetryFunc
	RetryIfFunc = retry.RetryIfFunc

	RetryableFunc                = retry.RetryableFunc
	RetryableFuncWithData[T any] = retry.RetryableFuncWithData[T]

	Timer = retry.Timer

	Error = retry.Error
)

// New creates a reusable Retrier.
//
// A Retrier is safe for concurrent use and can be reused across many retry
// operations. Reusing one avoids re-applying options and re-allocating the
// internal state on every call, which matters on hot paths.
//
// Retrier is not safe for concurrent use if any of its options mutate shared
// state, so options such as OnRetry must be free of side effects.
func New(opts ...Option) *Retrier {
	return retry.New(opts...)
}

// NewWithData creates a reusable RetrierWithData.
//
// It is the generic counterpart of New and returns the value produced by the
// retried function alongside its error.
func NewWithData[T any](opts ...Option) *RetrierWithData[T] {
	return retry.NewWithData[T](opts...)
}

// Do runs retryableFunc with the given options.
//
// It preserves the package-level convenience API of retry-go v4, which v5
// replaced with retry.New(...).Do(...).
//
// Each call allocates a fresh Retrier. On hot paths prefer New to build a
// Retrier once and reuse it.
func Do(retryableFunc RetryableFunc, opts ...Option) error {
	return retry.New(opts...).Do(retryableFunc)
}

// DoWithData runs retryableFunc with the given options and returns the value
// it produced.
//
// It is the generic counterpart of Do and replaces retry-go v4's
// package-level DoWithData.
func DoWithData[T any](retryableFunc RetryableFuncWithData[T], opts ...Option) (T, error) {
	return retry.NewWithData[T](opts...).Do(retryableFunc)
}
