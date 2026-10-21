package id

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func NanoID2(len int) string {
	if len <= 0 {
		idv, _ := gonanoid.New()
		return idv
	}
	idv, err := gonanoid.New(len)
	if err != nil {
		idv, _ = gonanoid.New()
	}
	return idv
}
