package registrar

import (
	"context"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

type RegisterFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	OpenSearchRegistrar(context.Context) RegisterFunc
}

type ServiceRegistrar = Client

type serviceRegistrar struct {
	*opensearchapi.Client
}

func NewServiceRegistrar(c *opensearchapi.Client) ServiceRegistrar {
	return serviceRegistrar{c}
}

type Func[S any] func(ServiceRegistrar, S)

type FuncWithErr[S any] func(ServiceRegistrar, S) error

type CtxFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxFuncWithErr[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToOpenSearch(context.Context, ServiceRegistrar) error
}

var _ Registrar = (*registrar[any])(nil)

func New[S any](s S, f Func[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewWithErr[S any](s S, f FuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxWithErr(s, cf)
}

func NewCtx[S any](s S, f CtxFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtxWithErr[S any](s S, f CtxFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

type registrar[S any] struct {
	s S
	f CtxFuncWithErr[S]
}

func (h *registrar[S]) RegisterToOpenSearch(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
