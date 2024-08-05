package realip

import (
	"github.com/blink-io/x/realip"
)

type options = realip.Options

type Option = realip.Option

func applyOption(ops ...Option) *options {
	opt := realip.DefaultOptions
	for _, o := range ops {
		o(opt)
	}
	return opt
}

var (
	WithHeaders = realip.WithHeaders

	WithRange = realip.WithRange
)
