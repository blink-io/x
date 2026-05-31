package testutil

import (
	"context"
	"log"
	"log/slog"

	"github.com/blink-io/x/internal/testdata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/stats"
)

func CreateGRPCClient(target string, secure bool, ops ...grpc.DialOption) *grpc.ClientConn {
	var creds credentials.TransportCredentials
	if secure {
		creds = credentials.NewTLS(testdata.GetClientTLSConfig())
	} else {
		creds = insecure.NewCredentials()
	}
	var newOps = make([]grpc.DialOption, 0)
	newOps = append(newOps, grpc.WithTransportCredentials(creds))
	if len(ops) > 0 {
		newOps = append(newOps, ops...)
	}
	c, err := grpc.Dial(target, newOps...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CreateGRPCServer(secure bool, ops ...grpc.ServerOption) *grpc.Server {
	var creds credentials.TransportCredentials
	if secure {
		creds = credentials.NewTLS(testdata.GetTLSConfig())
	} else {
		creds = insecure.NewCredentials()
	}

	var newOps = make([]grpc.ServerOption, 0)
	newOps = append(newOps, grpc.Creds(creds), grpc.StatsHandler(&statsHandler{}))
	if len(ops) > 0 {
		newOps = append(newOps, ops...)
	}
	gsrv := grpc.NewServer(newOps...)

	return gsrv
}

var _ stats.Handler = (*statsHandler)(nil)

type statsHandler struct {
}

func (s *statsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	slog.Info("Invoke [TagRPC]", "info", info)
	return ctx
}

func (s *statsHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	slog.Info("Invoke [HandleRPC]", "rpcStats", rpcStats)
}

func (s *statsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	slog.Info("Invoke [TagConn]", "remote addr", info.RemoteAddr)
	slog.Info("Invoke [TagConn]", "local addr", info.LocalAddr)
	return ctx
}

func (s *statsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	slog.Info("Invoke [HandleConn]", "connStats", connStats)
}
