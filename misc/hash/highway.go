package hash

import (
	"github.com/minio/highwayhash"
)

func Highway(key [128]byte, data []byte) [16]byte {
	return highwayhash.Sum128(data, key[:])
}
