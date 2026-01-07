package flatbuffers

import (
	"github.com/blink-io/x/encoding"

	flatbuffers "github.com/google/flatbuffers/go"
)

const (
	Name = "flatbuffers"
)

type codec struct {
	fc flatbuffers.FlatbuffersCodec
}

func New() encoding.Codec {
	return &codec{
		fc: flatbuffers.FlatbuffersCodec{},
	}
}

func (c *codec) Marshal(v any) ([]byte, error) {
	return c.fc.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v any) error {
	return c.fc.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
