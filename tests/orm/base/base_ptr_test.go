package base

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/ptr"
)

type model struct {
	name    string
	version string
	score   float64
	level   int
}

var mm = model{
	name:    "mm-name",
	version: "mm-version",
	score:   88.88,
	level:   6677,
}

func getMM() model {
	return mm
}

func getMMPtr() *model {
	return &mm
}

func TestPtr_1(t *testing.T) {
	mmptr := &mm
	mmptr1 := getMMPtr()

	name1 := "Hello"
	fmt.Printf("name1 ptr:%p\n", &name1)

	mmptr.name = name1

	fmt.Printf("mm ptr: %p\n", mmptr)
	fmt.Printf("mm name ptr before: %p\n", &mmptr.name)

	name2 := "你好"
	fmt.Printf("name2 ptr:%p\n", &name2)

	mmptr1.name = name2

	fmt.Printf("mm ptr1: %p\n", mmptr1)
	fmt.Printf("mm name ptr after: %p\n", &mmptr.name)
}

func TestPtr_2(t *testing.T) {
	mmptr := &mm
	mmptr.name = "after-after"

	fmt.Printf("mm ptr: %p\n", mmptr)
}

func TestPtr_3(t *testing.T) {
	mmptr := ptr.Of(mm)
	mmfptr := ptr.Of(getMM())

	fmt.Printf("mm ptr: %p\n", mmptr)
	fmt.Printf("mm fptr: %p\n", mmfptr)
}
