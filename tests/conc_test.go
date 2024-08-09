package tests

import (
	"fmt"
	"testing"

	"github.com/sourcegraph/conc"
)

func TestConc_1(t *testing.T) {
	var wg conc.WaitGroup

	wg.Go(func() {
		fmt.Println("Hello World")
	})

	wg.Wait()
}
