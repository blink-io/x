package realip

import (
	"net"
	"strings"
)

type Options struct {
	Headers        []string `json:"headers" yaml:"headers" toml:"headers" msgpack:"headers"`
	PrivateSubnets []Range  `json:"private_subnets" yaml:"private_subnets" toml:"private_subnets" msgpack:"private_subnets"`
}

// DefaultOptions is an `Options` value with some default headers and private subnets.
// See `Get` method.
var DefaultOptions = &Options{
	Headers: []string{
		HeaderXRealIP,
		HeaderXClientIP,
		HeaderXForwardedFor,
		HeaderCFConnectingIP,
	},
	PrivateSubnets: PrivateSubnets,
}

// addRange adds a private subnet to "opts".
// Should be called before any use of `Get`.
func (o *Options) addRange(start, end string) *Options {
	o.PrivateSubnets = append(o.PrivateSubnets, Range{
		Start: net.ParseIP(start),
		End:   net.ParseIP(end),
	})
	return o
}

// addHeaders adds a proxy remote address header to "opts".
// Should be called before any use of `Get`.
func (o *Options) addHeaders(headers ...string) *Options {
	o.Headers = append(o.Headers, headers...)
	return o
}

// Get extracts the real client's remote IP Address.
func (o *Options) Get(addrs ...string) string {
	for _, addr := range addrs {
		addrs := strings.Split(addr, ",")
		if ip, ok := GetIPAddress(addrs, o.PrivateSubnets); ok {
			return ip
		}
	}
	return ""
}

type Option func(*Options)

func WithHeaders(headers ...string) Option {
	return func(o *Options) {
		o.addHeaders(headers...)
	}
}

func WithRange(start, end string) Option {
	return func(o *Options) {
		o.addRange(start, end)
	}
}
