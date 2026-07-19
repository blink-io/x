package humah3

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	khttp3 "github.com/go-kratos/kratos/v3/transport/http"
)

// NewContext creates a new Huma context from an HTTP request and response.
func NewContext(op *huma.Operation, ctx khttp3.Context) huma.Context {
	return humamux.NewContext(op, ctx.Request(), ctx.Response())
}

var _ huma.Adapter = (*adapter)(nil)

type adapter struct {
	*khttp3.Server
	prefix string
}

func (a *adapter) Handle(op *huma.Operation, h func(huma.Context)) {
	rr := a.Route(a.prefix)
	rr.Handle(op.Method, op.Path, func(ctx khttp3.Context) error {
		h(NewContext(op, ctx))
		return nil
	})
}

func (a *adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Server.ServeHTTP(w, r)
}

func NewAdapter(srv *khttp3.Server, prefix string) huma.Adapter {
	return &adapter{Server: srv, prefix: prefix}
}

func New(srv *khttp3.Server, config huma.Config) huma.API {
	return NewWithPrefix(srv, "", config)
}

func NewWithPrefix(srv *khttp3.Server, prefix string, config huma.Config) huma.API {
	if len(config.Servers) == 0 {
		config.Servers = append(config.Servers, &huma.Server{
			URL: prefix,
		})
	}
	return huma.NewAPI(config, NewAdapter(srv, prefix))
}
