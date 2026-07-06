package registrar

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	kgrpc "github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/stretchr/testify/require"
)

func TestIsServerThen_1(t *testing.T) {
	ss := kgrpc.NewServer()

	rr := NewServiceRegistrar(ss)

	r := New[string]("test", func(r ServiceRegistrar, s string) {
		fmt.Println(r)
	})

	err := r.RegisterToGRPC(context.Background(), rr)
	require.NoError(t, err)

	IsServerThen(ss, func(s *grpc.Server) {
		fmt.Println("KKK")
	})
}
