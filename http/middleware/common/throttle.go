package common

import chimw "github.com/go-chi/chi/v5/middleware"

var (
	Throttle = chimw.Throttle

	ThrottleBacklog = chimw.ThrottleBacklog

	ThrottleWithOpts = chimw.ThrottleWithOpts
)
