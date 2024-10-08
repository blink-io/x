package gjson

import (
	"github.com/tidwall/gjson"
)

var (
	AddModifier      = gjson.AddModifier
	AppendJSONString = gjson.AppendJSONString
	Escape           = gjson.Escape
	ForEachLine      = gjson.ForEachLine
	ModifierExists   = gjson.ModifierExists
	Valid            = gjson.Valid
	ValidBytes       = gjson.ValidBytes
)

type (
	Result = gjson.Result
	Type   = gjson.Type
)
