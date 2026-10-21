package requestid

import (
	"context"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/misc/requestid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(ops ...Option) grpc.UnaryServerInterceptor {
	opts := applyOptions(ops...)
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		reqID := mdutil.SingleValueFromContext(ctx, opts.header)
		if len(reqID) == 0 {
			reqID = opts.generator()
		}
		// Append on top of the existing outgoing metadata instead of
		// rebuilding a map per request; rebuilt metadata silently dropped
		// entries added by upstream interceptors (auth, tracing).
		ctx = metadata.AppendToOutgoingContext(ctx, opts.header, reqID)
		resp, err = handler(ctx, req)
		return
	}
}

func StreamServerInterceptor(ops ...Option) grpc.StreamServerInterceptor {
	o := applyOptions(ops...)
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		reqID := mdutil.SingleValueFromContext(ss.Context(), o.header)
		if len(reqID) == 0 {
			reqID = o.generator()
		}
		wsCtx := requestid.NewContext(ss.Context(), reqID)
		ws := util.WrapServerStream(ss)
		ws.WrappedContext = metadata.AppendToOutgoingContext(wsCtx, o.header, reqID)
		ss = ws
		return handler(srv, ss)
	}
}
