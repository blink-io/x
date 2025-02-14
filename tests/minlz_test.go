package tests

import (
	"fmt"
	"testing"

	"github.com/minio/minlz"
)

func TestMinlz_1(t *testing.T) {
	l1 := minlz.LevelFastest
	fmt.Println(l1)
}
