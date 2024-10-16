package requestid

import (
	"github.com/google/uuid"
)

const (
	DefaultHeader = "X-Request-UserID"
)

type options struct {
	Header    string
	Generator func() string
}

var DefaultOptions = &options{
	Header:    DefaultHeader,
	Generator: defaultGenerator,
}

func setupOptions(c *options) *options {
	if c == nil {
		return DefaultOptions
	}
	if c.Header == "" {
		c.Header = DefaultHeader
	}
	if c.Generator == nil {
		c.Generator = defaultGenerator
	}
	return c
}

func defaultGenerator() string {
	return uuid.New().String()
}
