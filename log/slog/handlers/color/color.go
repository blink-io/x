package color

import (
	"gitlab.com/greyxor/slogor"
)

var (
	Err = slogor.Err
	New = slogor.NewHandler

	SetTimeFormat  = slogor.SetTimeFormat
	SetLevel       = slogor.SetLevel
	SetLevelStr    = slogor.SetLevelStr
	ShowSource     = slogor.ShowSource
	DisableColor   = slogor.DisableColor
	SetReplaceAttr = slogor.SetReplaceAttr
)

type (
	OptionFn = slogor.OptionFn
	Handler  = slogor.Handler

	MapOfLevel = slogor.MapOfLevel
)
