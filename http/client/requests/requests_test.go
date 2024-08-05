package requests

import (
	"context"
	"fmt"
	"testing"

	"github.com/blink-io/x/testing/tlsutil"
	"github.com/carlmjohnson/requests"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestReqs_1(t *testing.T) {
	var s string
	err := requests.URL("https://www.baidu.com").ToString(&s).Fetch(ctx)
	require.NoError(t, err)

	fmt.Println(s)
}

func TestHTTP3_1(t *testing.T) {
	tlsConfig, err := tlsutil.InsecureTLSConfig()
	require.NoError(t, err)

	h3c := HTTP3(tlsConfig)
	require.NotNil(t, h3c)
}
