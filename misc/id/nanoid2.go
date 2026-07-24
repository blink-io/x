package id

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func NanoID2(len int) string {
	idv, _ := gonanoid.New()
	return idv
}
