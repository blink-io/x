package hash

import (
	"encoding/binary"

	"github.com/zeebo/xxh3"
)

func XXH3(data []byte) []byte {
	return binary.BigEndian.AppendUint64(nil, xxh3.Hash(data))
}
