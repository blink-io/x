package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type (
	Context = khttp.Context

	// HandlerFunc defines a function to serve HTTP requests.
	HandlerFunc = khttp.HandlerFunc

	WalkRouteFunc = khttp.WalkRouteFunc

	FilterFunc = khttp.FilterFunc

	// Request type net/http.
	Request = http.Request

	// ResponseWriter type net/http.
	ResponseWriter = http.ResponseWriter

	// Flusher type net/http
	Flusher = http.Flusher

	DecodeRequestFunc = khttp.DecodeRequestFunc

	EncodeResponseFunc = khttp.EncodeResponseFunc

	EncodeErrorFunc = khttp.EncodeErrorFunc

	RouteInfo = khttp.RouteInfo
)

const SupportPackageIsVersion1 = khttp.SupportPackageIsVersion1

var (
	FilterChain = khttp.FilterChain

	NewRedirect = khttp.NewRedirect

	DefaultRequestVars = khttp.DefaultRequestVars

	DefaultRequestQuery = khttp.DefaultRequestQuery

	DefaultRequestDecoder = khttp.DefaultRequestDecoder

	DefaultResponseEncoder = khttp.DefaultResponseEncoder

	DefaultErrorEncoder = khttp.DefaultErrorEncoder

	CodecForRequest = khttp.CodecForRequest
)
