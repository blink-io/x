package loader

import "github.com/blink-io/x/i18n"

type options struct {
	bundle *i18n.Bundle
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := &options{}
	for _, o := range ops {
		o(opts)
	}
	if opts.bundle == nil {
		opts.bundle = i18n.Default()
	}
	return opts
}

func WithBundle(bundle *i18n.Bundle) Option {
	return func(o *options) {
		o.bundle = bundle
	}
}
