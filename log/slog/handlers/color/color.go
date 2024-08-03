package color

import (
	"gitlab.com/greyxor/slogor"
)

var (
	Err = slogor.Err
	New = slogor.NewHandler
)

type (
	Options      = slogor.Options
	Handler      = slogor.Handler
	GroupOrAttrs = slogor.GroupOrAttrs
)
