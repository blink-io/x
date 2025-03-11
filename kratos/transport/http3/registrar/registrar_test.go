package registrar

import (
	"context"
	"fmt"
	"testing"

	khttp3 "github.com/blink-io/kratos-transport/transport/http3"
	"github.com/stretchr/testify/require"
)

func TestIsServerThen_1(t *testing.T) {
	ss := khttp3.NewServer()

	rr := NewServiceRegistrar(ss)

	r := New[string]("test", func(r ServiceRegistrar, s string) {
		fmt.Println(r)
	})

	err := r.RegisterToHTTP3(context.Background(), rr)
	require.NoError(t, err)

	IsServerThen(ss, func(s *khttp3.Server) {
		fmt.Println("KKK")
	})
}
