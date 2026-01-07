package encoding

import (
	"time"
)

type Payload struct {
	Deadline time.Time      `json:"deadline" toml:"deadline" yaml:"deadline" msgpack:"deadline" cbor:"deadline"`
	Values   map[string]any `json:"values" toml:"values" yaml:"values" msgpack:"values" cbor:"values"`
}

// Codec is the interface for encoding/decoding session data to and from a byte
// slice for use by the session store.
type Codec interface {
	Encode(deadline time.Time, values map[string]any) ([]byte, error)
	Decode([]byte) (deadline time.Time, values map[string]any, err error)
}
