package xstrings

import (
	"fmt"
	"testing"

	"github.com/huandu/xstrings"
)

func TestXstrings_1(t *testing.T) {
	strs := []string{
		"no code",
		"DEV_ENV",
		"sys_log",
		"Unit_Code",
	}
	for _, str := range strs {
		fmt.Printf("%s ---> %s, %s, %s\n",
			str,
			xstrings.ToCamelCase(str),
			xstrings.ToKebabCase(str),
			xstrings.ToPascalCase(str),
		)
	}
}
