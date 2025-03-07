package conc

import (
	"github.com/negrel/conc"
)

type (
	Job[T any] = conc.Job[T]
	Routine    = conc.Routine
)
