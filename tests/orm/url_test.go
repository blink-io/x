package orm

import (
	"fmt"
	"net/url"
	"testing"
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
