package app

import (
	"strings"

	"github.com/avelino/slugify"
	"github.com/dimiro1/x/xtemplate"
)

func Uppercase() xtemplate.FuncMapping {
	return xtemplate.FuncMapping{
		Func: &xtemplate.Func{
			Name: "upper",
			Func: func(s string) string {
				return strings.ToUpper(s)
			},
		},
	}
}

func Slugify() xtemplate.FuncMapping {
	return xtemplate.FuncMapping{
		Func: &xtemplate.Func{
			Name: "slugify",
			Func: func(s string) string {
				return slugify.Slugify(s)
			},
		},
	}
}
