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

type Func[S any] func(context.Context, ServerHandler, S) error

type Registrar interface {
	RegisterToHTTP3(context.Context, ServerHandler) error
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

func (r registrar[S]) RegisterToHTTP3(ctx context.Context, sh ServerHandler) error {
	return r.f(ctx, sh, r.s)
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
