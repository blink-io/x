package tests

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/klauspost/compress/flate"
	"github.com/stretchr/testify/require"
)

func TestMy_Type_1(t *testing.T) {
	type Profile struct {
		Location  string
		CreatedAt time.Time
		Interval  time.Duration
	}
	type User struct {
		Name       string
		Age        int
		Profile    Profile
		ProfilePtr *Profile
	}

	var tt = time.Now()

	var u1 = User{
		Name: "ABC",
		Age:  23,
		Profile: Profile{
			Location:  "GZ, Guangdong",
			CreatedAt: tt,
		},
		ProfilePtr: &Profile{
			Location:  "NY, USA",
			CreatedAt: time.Now().UTC(),
		},
	}

	u2 := u1

	fmt.Printf("u1 name ptr: %p\n", &u1.Name)
	fmt.Printf("u2 name ptr: %p\n", &u2.Name)
	fmt.Printf("u1 profile ptr: %p\n", &u1.Profile)
	fmt.Printf("u2 profile ptr: %p\n", &u2.Profile)
}

const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`

func TestCompress_Flate_1(t *testing.T) {
	var b bytes.Buffer

	zw, err := flate.NewWriter(&b, flate.BestSpeed)
	require.NoError(t, err)

	_, err = io.Copy(zw, strings.NewReader(data))
	require.NoError(t, err)

	err = zw.Close()
	require.NoError(t, err)

	hstr := hex.EncodeToString(b.Bytes())
	fmt.Printf("%s, %d, raw: %d\n", hstr, len(hstr), len(data))

	zr := flate.NewReader(&b)
	_, err = io.Copy(os.Stdout, zr)
	require.NoError(t, err)

	err = zr.Close()
	require.NoError(t, err)
}
