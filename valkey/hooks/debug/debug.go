package debug

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeyhook"
)

type hook struct {
	logf func(string, ...any)
}

var _ valkeyhook.Hook = (*hook)(nil)

func New(ops ...Option) valkeyhook.Hook {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.logf == nil {
		h.logf = log.Printf
	}
	return h
}

func (h *hook) Do(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResult {
	h.logf("Redis CMD: [%s]", cmdstr(cmd.Commands()))
	return c.Do(ctx, cmd)
}

func (h *hook) DoMulti(c valkey.Client, ctx context.Context, multi ...valkey.Completed) []valkey.ValkeyResult {
	for _, m := range multi {
		h.logf("Valkey CMD: [%s]", cmdstr(m.Commands()))
	}
	rr := c.DoMulti(ctx, multi...)
	return rr
}

func (h *hook) DoCache(c valkey.Client, ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) valkey.ValkeyResult {
	h.logf("Valkey CMD: [%s]", cmdstr(cmd.Commands()))
	r := c.DoCache(ctx, cmd, ttl)
	return r
}

func (h *hook) DoMultiCache(c valkey.Client, ctx context.Context, multi ...valkey.CacheableTTL) []valkey.ValkeyResult {
	for _, m := range multi {
		h.logf("Valkey CMD: [%s]", cmdstr(m.Cmd.Commands()))
	}
	rr := c.DoMultiCache(ctx, multi...)
	return rr
}

func (h *hook) Receive(c valkey.Client, ctx context.Context, sub valkey.Completed, fn func(msg valkey.PubSubMessage)) error {
	h.logf("Valkey CMD: [%s]", cmdstr(sub.Commands()))
	err := c.Receive(ctx, sub, fn)
	return err
}

func (h *hook) DoStream(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream {
	h.logf("Valkey CMD: [%s]", cmdstr(cmd.Commands()))
	return c.DoStream(ctx, cmd)
}

func (h *hook) DoMultiStream(c valkey.Client, ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream {
	for _, m := range multi {
		h.logf("Valkey CMD: [%s]", cmdstr(m.Commands()))
	}
	rr := c.DoMultiStream(ctx, multi...)
	return rr
}

func cmdstr(cmds []string) string {
	return strings.Join(cmds, " ")
}
