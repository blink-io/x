package mdns

import (
	"context"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
)

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
	logger    klog.Logger
}

// Namespace with registry namespace.
func Namespace(ns string) Option {
	return func(o *options) { o.namespace = ns }
}

// RegisterTTL with register ttl.
func RegisterTTL(ttl time.Duration) Option {
	return func(o *options) { o.ttl = ttl }
}

func MaxRetry(num int) Option {
	return func(o *options) { o.maxRetry = num }
}
