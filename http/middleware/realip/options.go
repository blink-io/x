package realip

import "github.com/blink-io/x/misc/realip"

type options = realip.Options

type Option = realip.Option

func applyOption(ops ...Option) *options {
	// Deep-copy the shared defaults: options otherwise append into the
	// global DefaultOptions and leak into every other middleware instance.
	opt := &realip.Options{
		Headers:        append([]string(nil), realip.DefaultOptions.Headers...),
		PrivateSubnets: append([]realip.Range(nil), realip.DefaultOptions.PrivateSubnets...),
	}
	for _, o := range ops {
		o(opt)
	}
	return opt
}

var (
	WithHeaders = realip.WithHeaders

	WithRange = realip.WithRange
)
