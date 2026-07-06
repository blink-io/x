package registrar

import (
	"context"
	"net/http"

	khttp3 "github.com/blink-io/kratos-transport/transport/http3"
	khttp "github.com/go-kratos/kratos/v3/transport/http"
)

type (
	FilterFunc = khttp.FilterFunc
)

type RegisterFunc func(context.Context, ServerHandler) error

type WithRegistrar interface {
	HTTP3Registrar(context.Context) RegisterFunc
}

type ServerHandler interface {
	Handle(path string, h http.Handler)
	HandleFunc(path string, h http.HandlerFunc)
	HandleHeader(key, val string, h http.HandlerFunc)
	HandlePrefix(prefix string, h http.Handler)
}

var _ ServerHandler = (*khttp3.Server)(nil)

type serverHandler struct {
	*khttp3.Server
}

var _ ServerHandler = (*khttp3.Server)(nil)

func NewServerHandler(s *khttp3.Server) ServerHandler {
	return serverHandler{s}
}

type Func[S any] func(ServerHandler, S)

type FuncWithErr[S any] func(ServerHandler, S) error

type CtxFunc[S any] func(context.Context, ServerHandler, S)

type CtxFuncWithErr[S any] func(context.Context, ServerHandler, S) error

type Registrar interface {
	RegisterToHTTP3(context.Context, ServerHandler) error
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

func (h registrar[S]) RegisterToHTTP3(ctx context.Context, rr ServerHandler) error {
	return h.f(ctx, rr, h.s)
}

type Router interface {
	Route(prefix string, filters ...FilterFunc) *khttp3.Router
}

func IsRouterThen(s any, f func(Router)) {
	if v, ok := s.(Router); ok {
		f(v)
	}
}

func IsServerThen(s any, f func(*khttp3.Server)) {
	if v, ok := s.(*khttp3.Server); ok {
		f(v)
	}
}
