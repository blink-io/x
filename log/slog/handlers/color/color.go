package color

import (
	"gitlab.com/greyxor/slogor"
)

var (
	Err = slogor.Err
	New = slogor.NewHandler
)

type (
	OptionFn     = slogor.OptionFn
	Handler      = slogor.Handler
	GroupOrAttrs = slogor.GroupOrAttrs
)

var (
	SetTimeFormat = slogor.SetTimeFormat
	SetLevel      = slogor.SetLevel
	ShowSource    = slogor.ShowSource
	DisableColor  = slogor.DisableColor
)
