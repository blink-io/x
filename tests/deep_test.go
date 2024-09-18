package tests

import (
	"testing"

	"github.com/go-test/deep"
)

type T struct {
	Name    string
	Numbers []float64
}

func TestDeepEqual(t *testing.T) {
	// Can you spot the difference?
	t1 := T{
		Name:    "Isabella",
		Numbers: []float64{1.13459, 2.29343, 3.010100010},
	}
	t2 := T{
		Name:    "Isabella",
		Numbers: []float64{1.13459, 2.29843, 3.010100010},
	}

	if diff := deep.Equal(t1, t2); diff != nil {
		t.Error(diff)
	}
}
