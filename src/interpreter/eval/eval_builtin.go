package eval

import (
	"github.com/ape-lang/ape/src/data"
)

var builtins = map[string]*data.Builtin{
	"len":   data.GetBuiltinDef("len"),
	"head":  data.GetBuiltinDef("head"),
	"tail":  data.GetBuiltinDef("tail"),
	"last":  data.GetBuiltinDef("last"),
	"push":  data.GetBuiltinDef("push"),
	"print": data.GetBuiltinDef("print"),
}

func evalBuiltin(value string) (*data.Builtin, bool) {
	result, ok := builtins[value]
	return result, ok
}
