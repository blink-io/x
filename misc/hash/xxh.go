package hash

import (
	"encoding/binary"

	"github.com/cespare/xxhash/v2"
)

func XXH(data []byte) []byte {
	return binary.BigEndian.AppendUint64(nil, xxhash.Sum64(data))
}
