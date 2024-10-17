package registrar

import (
	"context"
	"fmt"
	"testing"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/require"
)

func TestIsServerThen_1(t *testing.T) {
	ss := khttp.NewServer()

	rr := NewServiceRegistrar(ss)

	r := New[string]("test", func(r ServiceRegistrar, s string) {
		fmt.Println(r)
	})

	err := r.RegisterToHTTP(context.Background(), rr)
	require.NoError(t, err)

	IsServerThen(ss, func(s *khttp.Server) {
		fmt.Println("KKK")
	})
}
