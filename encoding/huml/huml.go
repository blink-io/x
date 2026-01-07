package huml

import (
	"github.com/blink-io/x/encoding"
	"github.com/huml-lang/go-huml"
)

const (
	Name = "huml"
)

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v any) ([]byte, error) {
	return huml.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v any) error {
	return huml.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
