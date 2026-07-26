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

type Func[S any] func(context.Context, ServerHandler, S) error

type Registrar interface {
	RegisterToHTTP(context.Context, ServerHandler) error
}

type registrar[S any] struct {
	s S
	f Func[S]
}

var _ Registrar = (*registrar[any])(nil)

func NewRegistrar[S any](s S, f Func[S]) Registrar {
	rr := &registrar[S]{
		s: s,
		f: f,
	}
	return rr
}

func (r registrar[S]) RegisterToHTTP(ctx context.Context, sh ServerHandler) error {
	return r.f(ctx, sh, r.s)
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
