package registrar

import (
	"context"
	"fmt"
	"testing"

	khttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/stretchr/testify/require"
)

func TestIsServerThen_1(t *testing.T) {
	ss := khttp.NewServer()

	rr := NewServerHandler(ss)

	r := NewRegistrar[string]("test", func(ctx context.Context, sh ServerHandler, s string) error {
		fmt.Println(sh)
		return nil
	})

	err := r.RegisterToHTTP(context.Background(), rr)
	require.NoError(t, err)

	IsServerThen(ss, func(s *khttp.Server) {
		fmt.Println("KKK")
	})
}
