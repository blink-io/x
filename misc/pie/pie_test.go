package pie

import (
	"testing"
)

func TestPie(t *testing.T) {
	ss := []string{
		"1", "2", "3",
	}
	_ = Contains(ss, "2")
}
