package pb

import (
	"testing"

	"github.com/go-kratos/kratos/v2/encoding"
	_ "github.com/go-kratos/kratos/v2/encoding/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
)

func TestPB_Message_1(t *testing.T) {
	req := &grpc_reflection_v1.ExtensionRequest{
		ContainingType:  "ABCDEFG-ok,yes，中文",
		ExtensionNumber: int32(1886),
	}
	codec := encoding.GetCodec("proto")
	require.NotNil(t, codec)

	data, err := codec.Marshal(req)
	require.NoError(t, err)

	var req2 = new(grpc_reflection_v1.ExtensionRequest)
	err = codec.Unmarshal(data, req2)
	require.NoError(t, err)
}
