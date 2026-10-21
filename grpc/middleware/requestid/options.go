package requestid

import (
	"strings"

	"github.com/google/uuid"

	"github.com/blink-io/x/misc/requestid"
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
	// gRPC metadata keys are lowercase; normalize once at construction
	// instead of per intercepted call.
	opt.header = strings.ToLower(opt.header)
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
