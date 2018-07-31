package app

import (
	"strings"

	"github.com/avelino/slugify"
	"github.com/dimiro1/x/xtemplate"
)

func Uppercase() xtemplate.FuncMapping {
	return xtemplate.NewFuncMapping("upper", strings.ToUpper)
}

func Slugify() xtemplate.FuncMapping {
	return xtemplate.NewFuncMapping("slugify", slugify.Slugify)
}
