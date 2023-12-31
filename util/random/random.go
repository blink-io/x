package random

import (
	"math/rand"
	"strings"
	"time"
)

type (
	Random struct {
		src rand.Source
	}
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
)

var (
	global = New()
)

func New() *Random {
	r := new(Random)
	r.src = rand.NewSource(time.Now().UnixNano())
	return r
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.src.Int63()%int64(len(charset))]
	}
	return string(b)
}

func String(length uint8, charsets ...string) string {
	return global.String(length, charsets...)
}
