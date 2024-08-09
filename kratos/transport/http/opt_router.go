package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

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
