package id

import (
	"sync"

	"github.com/jaevor/go-nanoid"
)

type nanoGen struct {
	once sync.Once
	fn   func() string
}

var (
	nanoGens     sync.Map // int -> *nanoGen
	nanoOnce     sync.Once
	nanoFallback func() string
)

func NanoID(length int) string {
	if length <= 0 {
		length = 21
	} else if length > 255 {
		length = 255
	}
	e, _ := nanoGens.LoadOrStore(length, &nanoGen{})
	g := e.(*nanoGen)
	g.once.Do(func() {
		if f, err := nanoid.Standard(length); err == nil {
			g.fn = f
		}
	})
	if g.fn == nil {
		nanoOnce.Do(func() {
			if f, err := nanoid.Standard(21); err == nil {
				nanoFallback = f
			}
		})
		return nanoFallback()
	}
	return g.fn()
}
