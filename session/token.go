package session

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/xid"
)

func SecureTokenGen() (string, error) {
	return defaultTokenGen()
}

func UUIDTokenGen() (string, error) {
	return uuid.NewString(), nil
}

func ShortUUIDTokenGen() (string, error) {
	return shortuuid.New(), nil
}

func XIDTokenGen() (string, error) {
	return xid.New().String(), nil
}
