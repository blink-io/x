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

type testBun struct {
	Name    string
	Version string
}

func (t testBun) Ptr() *testBun {
	return &t
}

func (t testBun) Change() {
	tptr := &t
	fmt.Printf("in Change t ptr: %p\n", tptr)
	tptr.Name = "name is changed"
}

func (t *testBun) Update() {
	fmt.Printf("in Update t ptr: %p\n", t)
	t.Name = "name is changed"
}

func Test(t *testing.T) {
	tk := testBun{
		Name:    "hello",
		Version: "ver",
	}

	fmt.Printf("outer t ptr: %p\n", &tk)

	fmt.Println(litter.Sdump(tk))

	tk.Change()

	fmt.Println(litter.Sdump(tk))

	tk.Update()

	fmt.Println(litter.Sdump(tk))
}
