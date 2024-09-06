package redigo

import "github.com/gomodule/redigo/redis"

type (
	Conn = redis.Conn
)

var (
	Dial        = redis.Dial
	DialContext = redis.DialContext
	DialURL     = redis.DialURL
)
