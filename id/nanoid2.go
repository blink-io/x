package id

import (
	"github.com/matoous/go-nanoid"
)

func NanoID2(len int) string {
	idv, _ := gonanoid.Nanoid(len)
	return idv
}
