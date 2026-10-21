package realip

import (
	"context"
	"net"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"github.com/blink-io/x/misc/realip"
)

func getFromGRPC(ctx context.Context, o *options) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	for _, header := range o.Headers {
		mdvals := md.Get(header)
		for _, mdval := range mdvals {
			if mdval == "" {
				continue
			}
			// Single-value headers (the common case) skip the slice
			// allocation of strings.Split entirely.
			addrs := []string{mdval}
			if strings.Contains(mdval, ",") {
				addrs = strings.Split(mdval, ",")
			}
			if ip, ok := realip.GetIPAddress(addrs, o.PrivateSubnets); ok {
				return ip
			}
		}
	}

	addr := getGRPCPeerAddr(ctx)
	if addr != "" {
		if ip, _, err := net.SplitHostPort(addr); err == nil {
			return ip
		}
	}

	return addr
}

// getGRPCPeerAddr get peer addr
func getGRPCPeerAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}
