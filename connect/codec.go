package connect

import (
	"connectrpc.com/connect"
	gojson "github.com/goccy/go-json"
	"github.com/shamaton/msgpack/v3"
)

const (
	codecNameJSON    = "json"
	codecNameMsgPack = "msgpack"
)

var _ connect.Codec = (*jsonCodec)(nil)

type jsonCodec struct{}

func (c *jsonCodec) Name() string {
	return codecNameJSON
}

func (c *jsonCodec) Marshal(v any) ([]byte, error) {
	return gojson.Marshal(v)
}

func (c *jsonCodec) Unmarshal(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}

var _ connect.Codec = (*jsonCodec)(nil)

type msgpackCodec struct{}

func (c *msgpackCodec) Name() string {
	return codecNameMsgPack
}

func (c *msgpackCodec) Marshal(v any) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c *msgpackCodec) Unmarshal(data []byte, v any) error {
	return msgpack.Unmarshal(data, v)
}
