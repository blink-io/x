package orm

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/sanity-io/litter"
)

func TestURL_1(t *testing.T) {
	u := &url.URL{
		Scheme:   "postgres",
		Host:     "192.168.50.88:5432",
		Path:     "orm_demo",
		User:     url.UserPassword("blink", "888asdf!#%"),
		RawQuery: "sslmode=disable",
	}

	fmt.Println(u)
	fmt.Println(u.RequestURI())
}

func TestSlicePtr_1(t *testing.T) {
	ss := []Tag{
		{
			Name: "a",
			Code: "a",
		},
		{
			Name: "b",
			Code: "b",
		},
		{
			Name: "c",
			Code: "c",
		},
	}

	fmt.Println(litter.Sdump(ss))

	p1ptr := &ss[1]
	p1ptr.Name = "bbbbb"
	p1ptr.Code = "bbbbb"

	fmt.Println(litter.Sdump(ss))
}
