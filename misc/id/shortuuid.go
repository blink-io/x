package id

import (
	"github.com/lithammer/shortuuid/v5"
)

func ShortUUID() string {
	return shortuuid.New()
}
