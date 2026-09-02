package tests

import (
	"testing"

	"github.com/guregu/null/v6"
)

func TestNull_1(t *testing.T) {
	nstr1 := null.StringFrom("hlello")
	println("IsZero ", nstr1.IsZero())
	println("Valid ", nstr1.Valid)

	nstr2 := null.StringFromPtr(nil)
	println("IsZero ", nstr2.IsZero())
	println("Valid ", nstr2.Valid)
}
