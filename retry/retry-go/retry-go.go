package retry_go

import "github.com/avast/retry-go/v4"

var (
	BackOffDelay = retry.BackOffDelay
	Do           = retry.Do
)

type (
	Config        = retry.Config
	Option        = retry.Option
	DelayTypeFunc = retry.DelayTypeFunc
	OnRetryFunc   = retry.OnRetryFunc
	RetryIfFunc   = retry.RetryIfFunc
	RetryableFunc = retry.RetryableFunc
	Timer         = retry.Timer
)
