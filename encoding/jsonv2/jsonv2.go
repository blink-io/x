//go:build !goexperiment.jsonv2

package jsonv2

import (
	"github.com/go-json-experiment/json"
)

type (
	Marshaler = json.Marshaler
)
