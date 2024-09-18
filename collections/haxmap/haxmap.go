package haxmap

import (
	"unsafe"

	"github.com/alphadose/haxmap"
	"golang.org/x/exp/constraints"
)

type (
	hashable interface {
		constraints.Integer | constraints.Float | constraints.Complex | ~string | uintptr | ~unsafe.Pointer
	}
)

func New[K hashable, V any]() *haxmap.Map[K, V] {
	return haxmap.New[K, V]()
}
