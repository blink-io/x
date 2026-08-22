package redis

import (
	redislock "github.com/go-co-op/gocron-redis-lock/v2"
)

var NewRedisLocker = redislock.NewRedisLocker
