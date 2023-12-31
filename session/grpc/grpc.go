package grpc

import (
	"context"
	"strings"
	"time"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/session"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const DefaultHeader = "X-Auth-Token"

const DefaultExpirySuffix = "-expiry"

type SessionHandler struct {
	manager session.Manager

	header string

	exposeExpiry bool

	expirySuffix string

	timeFormat string
}

func NewSessionHandler(ops ...Option) *SessionHandler {
	sh := &SessionHandler{
		manager:      session.NewManager(),
		timeFormat:   time.RFC3339Nano,
		expirySuffix: DefaultExpirySuffix,
	}
	for _, o := range ops {
		o(sh)
	}
	return sh
}

func UnaryServerInterceptor(sh *SessionHandler) grpc.UnaryServerInterceptor {
	return sh.UnaryServerInterceptor
}

func StreamServerInterceptor(sh *SessionHandler) grpc.StreamServerInterceptor {
	return sh.StreamServerInterceptor
}

func (sh *SessionHandler) UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	sm := sh.manager
	ctx = session.NewContext(ctx, sm)
	header := strings.ToLower(sh.header)
	token := mdutil.SingleValueFromContext(ctx, header)

	ctx, err := sm.Load(ctx, token)
	if err != nil {
		return nil, err
	}

	xres, xerr := handler(ctx, req)
	if xerr != nil {
		return nil, xerr
	}

	if err := sh.commitAndWriteSession(ctx); err != nil {
		return nil, err
	}

	return xres, xerr
}

func (sh *SessionHandler) StreamServerInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	sm := sh.manager
	ctx := session.NewContext(ss.Context(), sm)
	wss := util.WrapServerStream(ss)
	wss.WrappedContext = ctx

	//NOTICE In MD, all keys are lower characters.
	header := strings.ToLower(sh.header)
	token := mdutil.SingleValueFromContext(ctx, header)

	ctx, err := sm.Load(ctx, token)
	if err != nil {
		return err
	}

	if xerr := handler(srv, ss); xerr != nil {
		return xerr
	}

	if err := sh.commitAndWriteSession(ctx); err != nil {
		return err
	}

	return nil
}

func (sh *SessionHandler) commitAndWriteSession(ctx context.Context) error {
	headerKey := strings.ToLower(sh.header)
	expiryKey := headerKey + "-expiry"
	sm := sh.manager
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}

	switch sm.Status(ctx) {
	case session.Modified:
		token, expiry, err := sm.Commit(ctx)
		expiryValue := expiry.Format(sh.timeFormat)
		if err != nil {
			return err
		}
		md.Set(headerKey, token)
		if sh.exposeExpiry {
			md.Set(expiryKey, expiryValue)
		}
	case session.Destroyed:
		md.Delete(headerKey)
		if sh.exposeExpiry {
			md.Delete(expiryKey)
		}
	}

	return multierr.Combine(
		grpc.SendHeader(ctx, md),
		grpc.SetTrailer(ctx, md),
	)
}
