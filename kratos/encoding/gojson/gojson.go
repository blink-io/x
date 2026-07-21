package gojson

import (
	"github.com/blink-io/x/encoding/gojson"

	"github.com/go-kratos/kratos/v3/encoding"
)

func New() encoding.Codec {
	return gojson.New()
}
