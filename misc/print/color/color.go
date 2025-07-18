package color

import (
	"github.com/fatih/color"
)

type (
	Color      = color.Color
	Attributed = color.Attribute
)

const (
	FgWhite   = color.FgWhite
	FgBlack   = color.FgBlack
	FgRed     = color.FgRed
	FgGreen   = color.FgGreen
	FgYellow  = color.FgYellow
	FgBlue    = color.FgBlue
	FgMagenta = color.FgMagenta
	FgCyan    = color.FgCyan

	BgWhite  = color.BgWhite
	BgBlack  = color.BgBlack
	BgRed    = color.BgRed
	BgGreen  = color.BgGreen
	BgYellow = color.BgYellow
	BgBlue   = color.BgBlue
)

var (
	New = color.New
)
