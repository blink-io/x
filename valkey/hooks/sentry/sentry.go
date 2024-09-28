package sentry

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeyhook"
)

type hook struct {
	hub *sentry.Hub
}

var _ valkeyhook.Hook = (*hook)(nil)

func New(ops ...Option) (valkeyhook.Hook, error) {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	return h, nil
}

func (h *hook) Do(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResult {
	r := c.Do(ctx, cmd)
	h.handleSingleError(r)
	return r
}

func (h *hook) DoMulti(c valkey.Client, ctx context.Context, multi ...valkey.Completed) []valkey.ValkeyResult {
	ra := c.DoMulti(ctx, multi...)
	for _, r := range ra {
		if err := r.Error(); err != nil {
			h.handleSingleError(r)
		}
	}
	return ra
}

func (h *hook) DoCache(c valkey.Client, ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) valkey.ValkeyResult {
	r := c.DoCache(ctx, cmd, ttl)
	h.handleSingleError(r)
	return r
}

func (h *hook) DoMultiCache(c valkey.Client, ctx context.Context, multi ...valkey.CacheableTTL) []valkey.ValkeyResult {
	return c.DoMultiCache(ctx, multi...)
}

func (h *hook) Receive(c valkey.Client, ctx context.Context, sub valkey.Completed, fn func(msg valkey.PubSubMessage)) error {
	if err := c.Receive(ctx, sub, fn); err != nil {
		h.hub.CaptureException(err)
		return err
	}
	return nil
}

func (h *hook) DoStream(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream {
	ra := c.DoStream(ctx, cmd)
	for ra.HasNext() {
		if err := ra.Error(); err != nil {
			h.handleStreamError(ra)
		}
	}
	return ra
}

func (h *hook) DoMultiStream(c valkey.Client, ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream {
	ra := c.DoMultiStream(ctx, multi...)
	for ra.HasNext() {
		if err := ra.Error(); err != nil {
			h.handleStreamError(ra)
		}
	}
	return ra
}

func (h *hook) handleStreamError(r valkey.MultiValkeyResultStream) {
	if err := r.Error(); err != nil {
		h.hub.CaptureException(err)
	}
}

func (h *hook) handleSingleError(r valkey.ValkeyResult) {
	if err := r.Error(); err != nil {
		h.hub.CaptureException(err)
	}
}
