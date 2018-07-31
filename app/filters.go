package app

import (
	"strings"

	"github.com/avelino/slugify"
	"github.com/dimiro1/x/xtemplate"
)

func Uppercase() xtemplate.FuncMapping {
	return xtemplate.Register("upper", strings.ToUpper)
}

func Slugify() xtemplate.FuncMapping {
	return xtemplate.Register("slugify", slugify.Slugify)
}
