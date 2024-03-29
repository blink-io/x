package dbr

import (
	"github.com/gocraft/dbr/v2"
)

type options struct {
	er dbr.EventReceiver
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithEventReceiver(er dbr.EventReceiver) Option {
	return func(o *options) {
		o.er = er
	}
}
