package registrar

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	khttp3 "github.com/tx7do/kratos-transport/transport/http3"
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
