package requestid

import (
	"github.com/blink-io/x/requestid"
	"github.com/google/uuid"
)

type options struct {
	header    string
	generator func() string
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opt := &options{
		header:    requestid.DefaultHeader,
		generator: func() string { return uuid.NewString() },
	}
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func Generator(generator func() string) Option {
	return func(o *options) {
		if generator != nil {
			o.generator = generator
		}
	}
}

func Header(header string) Option {
	return func(o *options) {
		if len(header) > 0 {
			o.header = header
		}
	}
}
