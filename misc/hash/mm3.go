package hash

import (
	"encoding/binary"

	"github.com/twmb/murmur3"
)

// MM3 is short for murmur3
func MM3(data []byte) []byte {
	h1, h2 := murmur3.Sum128(data)
	out := binary.BigEndian.AppendUint64(nil, h1)
	return binary.BigEndian.AppendUint64(out, h2)
}
