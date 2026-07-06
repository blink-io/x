package http

import (
	"context"
	"fmt"
	"net/http"

	khttp "github.com/go-kratos/kratos/v3/transport/http"
)

const (
	MethodAny = "*"
)

type keyValue struct {
	key   string
	value string
}

type doOptions struct {
	//method      string
	statusCode  int
	ahs         []*keyValue
	shs         map[string]string
	skipQuery   bool
	skipVars    bool
	skipReqBody bool
	skipResBody bool
}

type (
	Req = any

	Res = any

	Func[Request Req, Response Res] func(context.Context, *Request) (*Response, error)
)

type DoOption func(*doOptions)

func applyDoOptions(ops ...DoOption) *doOptions {
	opt := &doOptions{
		ahs: make([]*keyValue, 0),
		shs: make(map[string]string),
	}
	for _, o := range ops {
		o(opt)
	}
	//if len(opt.method) == 0 {
	//	opt.method = http.MethodGet
	//}
	if opt.statusCode == 0 {
		opt.statusCode = http.StatusOK
	}
	return opt
}

func ApplyStatusCode(statusCode int) DoOption {
	return func(o *doOptions) {
		o.statusCode = statusCode
	}
}

func ApplyAddHeader(key string, value string) DoOption {
	return func(o *doOptions) {
		o.ahs = append(o.ahs, &keyValue{key, value})
	}
}

func ApplySetHeader(key string, value string) DoOption {
	return func(o *doOptions) {
		o.shs[key] = value
	}
}

func ApplySkipVars() DoOption {
	return func(o *doOptions) {
		o.skipVars = true
	}
}

func ApplySkipQuery() DoOption {
	return func(o *doOptions) {
		o.skipQuery = true
	}
}

func ApplySkipReqBody() DoOption {
	return func(o *doOptions) {
		o.skipReqBody = true
	}
}

func ApplySkipResBody() DoOption {
	return func(o *doOptions) {
		o.skipResBody = true
	}
}

func Do[Request Req, Response Res](
	method string,
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return func(kctx khttp.Context) error {
		opts := applyDoOptions(ops...)
		var in Request

		if !opts.skipQuery {
			if err := kctx.BindQuery(&in); err != nil {
				return err
			}

		}
		if !opts.skipVars {
			if err := kctx.BindVars(&in); err != nil {
				return err
			}
		}

		switch method {
		case http.MethodPost,
			http.MethodPut,
			// HTTP DELETE Maybe has payload
			// https://developer.mozilla.org/docs/Web/HTTP/Methods/DELETE
			http.MethodDelete,
			http.MethodPatch,
			MethodAny:
			if !opts.skipReqBody {
				if err := kctx.Bind(&in); err != nil {
					return err
				}
			}
			break
		default:
		}

		khttp.SetOperation(kctx, operation)
		mwh := kctx.Middleware(func(ctx context.Context, req any) (any, error) {
			reqValue, ok := req.(*Request)
			if !ok {
				return nil, fmt.Errorf("unexpected request type %T", req)
			}
			return handle(kctx, reqValue)
		})
		out, err := mwh(kctx, &in)
		if err != nil {
			return err
		}

		for _, i := range opts.ahs {
			kctx.Header().Add(i.key, i.value)
		}
		for k, v := range opts.shs {
			kctx.Header().Set(k, v)
		}

		if opts.skipResBody {
			return kctx.Result(opts.statusCode, nil)
		} else {
			reply, ok := out.(*Response)
			if !ok {
				return fmt.Errorf("unexpected response type %T", out)
			}
			return kctx.Result(opts.statusCode, reply)
		}
	}
}

func GET[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodGet, operation, handle, ops...)
}

func POST[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPost, operation, handle, ops...)
}

func PUT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPut, operation, handle, ops...)
}

func PATCH[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPatch, operation, handle, ops...)
}

func DELETE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodDelete, operation, handle, ops...)
}

func OPTIONS[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodOptions, operation, handle, ops...)
}

func CONNECT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodConnect, operation, handle, ops...)
}

func TRACE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodTrace, operation, handle, ops...)
}

func HEAD[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodHead, operation, handle, ops...)
}

func Any[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](MethodAny, operation, handle, ops...)
}

func (h Func[Request, Response]) Do(method, operation string, ops ...DoOption) khttp.HandlerFunc {
	return Do[Request, Response](method, operation, h, ops...)
}

func (h Func[Request, Response]) GET(operation string, ops ...DoOption) khttp.HandlerFunc {
	return GET[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) POST(operation string, ops ...DoOption) khttp.HandlerFunc {
	return POST[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) PUT(operation string, ops ...DoOption) khttp.HandlerFunc {
	return PUT[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) PATCH(operation string, ops ...DoOption) khttp.HandlerFunc {
	return PATCH[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) CONNECT(operation string, ops ...DoOption) khttp.HandlerFunc {
	return CONNECT[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) DELETE(operation string, ops ...DoOption) khttp.HandlerFunc {
	return DELETE[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) OPTIONS(operation string, ops ...DoOption) khttp.HandlerFunc {
	return OPTIONS[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) TRACE(operation string, ops ...DoOption) khttp.HandlerFunc {
	return TRACE[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) HEAD(operation string, ops ...DoOption) khttp.HandlerFunc {
	return HEAD[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) Any(operation string, ops ...DoOption) khttp.HandlerFunc {
	return Any[Request, Response](operation, h, ops...)
}

type RouterHandle interface {
	Handle(method, path string, h khttp.HandlerFunc, filters ...khttp.FilterFunc)
}

type RouterAction[Request Req, Response Res] func(path string, operation string, handle Func[Request, Response], ops ...DoOption)

func Route[Request Req, Response Res](rh RouterHandle, method string) RouterAction[Request, Response] {
	return func(path string, operation string, handle Func[Request, Response], ops ...DoOption) {
		rh.Handle(method, path, Do(method, operation, handle, ops...))
	}
}

func RouteCONNECT[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodConnect)
}

func RouteDELETE[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodDelete)
}

func RouteGET[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodGet)
}

func RouteHEAD[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodHead)
}

func RouteOPTIONS[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodOptions)
}

func RoutePATCH[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodPatch)
}

func RoutePOST[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodPost)
}

func RoutePUT[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodPut)
}

func RouteTRACE[Request Req, Response Res](r RouterHandle) RouterAction[Request, Response] {
	return Route[Request, Response](r, http.MethodTrace)
}
