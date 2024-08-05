package requestid

import (
	"context"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/requestid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(ops ...Option) grpc.UnaryServerInterceptor {
	opts := applyOptions(ops...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		reqID := mdutil.SingleValueFromContext(ctx, opts.header)
		if len(reqID) == 0 {
			reqID = opts.generator()
		}
		outMD := metadata.New(map[string]string{
			opts.header: reqID,
		})
		resp, err = handler(metadata.NewOutgoingContext(ctx, outMD), req)
		return
	}
}

func StreamServerInterceptor(ops ...Option) grpc.StreamServerInterceptor {
	o := applyOptions(ops...)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		reqID := mdutil.SingleValueFromContext(ss.Context(), o.header)
		if len(reqID) == 0 {
			reqID = o.generator()
		}
		outMD := metadata.New(map[string]string{
			o.header: reqID,
		})
		wsCtx := requestid.NewContext(ss.Context(), reqID)
		ws := util.WrapServerStream(ss)
		ws.WrappedContext = metadata.NewOutgoingContext(wsCtx, outMD)
		ss = ws
		return handler(srv, ss)
	}
}
