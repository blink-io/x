package registrar

import (
	"context"
	"net/http"

	khttp "github.com/go-kratos/kratos/v3/transport/http"
)

type RegisterFunc func(context.Context, ServerHandler) error

type WithRegistrar interface {
	HTTPRegistrar(context.Context) RegisterFunc
}

type ServerHandler interface {
	Handle(path string, h http.Handler)
	HandleFunc(path string, h http.HandlerFunc)
	HandleHeader(key, val string, h http.HandlerFunc)
	HandlePrefix(prefix string, h http.Handler)
}

var _ ServerHandler = (*khttp.Server)(nil)

type serverHandler struct {
	*khttp.Server
}

func NewServerHandler(s *khttp.Server) ServerHandler {
	return serverHandler{Server: s}
}

type Func[S any] func(ServerHandler, S)

type FuncWithErr[S any] func(ServerHandler, S) error

type CtxFunc[S any] func(context.Context, ServerHandler, S)

type CtxFuncWithErr[S any] func(context.Context, ServerHandler, S) error

type Registrar interface {
	RegisterToHTTP(context.Context, ServerHandler) error
}

type registrar[S any] struct {
	s S
	f CtxFuncWithErr[S]
}

var _ Registrar = (*registrar[any])(nil)

func New[S any](s S, f Func[S]) Registrar {
	cf := func(ctx context.Context, r ServerHandler, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtx[S any](s S, f CtxFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServerHandler, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

// NewWithErr creates a registrar with returning error.
func NewWithErr[S any](s S, f FuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServerHandler, s S) error {
		return f(r, s)
	}
	return NewCtxWithErr(s, cf)
}

// NewCtxWithErr creates a registrar with a context parameter and returning error.
func NewCtxWithErr[S any](s S, f CtxFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

func (h registrar[S]) RegisterToHTTP(ctx context.Context, rr ServerHandler) error {
	return h.f(ctx, rr, h.s)
}

type Router interface {
	Route(prefix string, filters ...khttp.FilterFunc) *khttp.Router
}

func IsRouterThen(s any, f func(Router)) {
	if v, ok := s.(Router); ok {
		f(v)
	}
}

func IsServerThen(s any, f func(*khttp.Server)) {
	if v, ok := s.(*khttp.Server); ok {
		f(v)
	}
}
