package tests

import (
	"fmt"
	"testing"
	"unique"

	"github.com/blink-io/x/ptr"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFull_1(t *testing.T) {
	t.Run("before", func(t *testing.T) {
		fmt.Println("run before")
	})
	fmt.Println("Run internally")
	t.Run("after", func(t *testing.T) {
		fmt.Println("run after")
	})
}

func TestUnique_1(t *testing.T) {
	type Args struct {
		Line   string
		Number int
	}

	a := Args{
		Line:   "System.out.println",
		Number: 188,
	}

	ax := Args{
		Line:   "System.out.println",
		Number: 188,
	}

	aptr := &Args{
		Line:   "What is the system design",
		Number: 88,
	}
	bptr := &Args{
		Line:   "What is the system design",
		Number: 88,
	}

	h1 := unique.Make(a)

	a1 := h1.Value()
	a2 := h1.Value()

	a3 := a
	a4 := a

	h2 := unique.Make(bptr)

	b1 := h2.Value()
	b2 := h2.Value()

	fmt.Printf("a1:%p\na2:%p\n", &a1, &a2)
	fmt.Printf("a3:%p\na4:%p\n", &a3, &a4)
	fmt.Println("----------------------------")
	fmt.Printf("b1:%p\nb2:%p\n", b1, b2)
	fmt.Println("----------------------------")

	assert.True(t, a1 == a2)
	assert.True(t, a3 == a4)
	assert.True(t, a3 == ax)
	assert.True(t, *aptr == *bptr)
	assert.True(t, b1 == b2)

	assert.False(t, aptr == bptr)
}

func TestSlice_1(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}

	arr := []user{
		{
			Name: "John",
			Age:  88,
		},
		{
			Name: "mary",
			Age:  33,
		},
		{
			Name: "nick",
			Age:  21,
		},
	}

	arrPtr := ptr.SliceOfPtrs(arr...)
	require.NotNil(t, arrPtr)

	fmt.Println("before", litter.Sdump(arr))

	a1ptr := &arr[1]
	a1ptr.Name = "Super Mary"
	a1ptr.Age = 77

	a2dir := arr[2]
	a2dir.Name = "Super Nick"

	fmt.Println("after", litter.Sdump(arr))
}
