package gojson

import (
	"github.com/blink-io/x/encoding"

	gojson "github.com/goccy/go-json"
)

const (
	Name = "gojson"
)

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v any) ([]byte, error) {
	return gojson.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
