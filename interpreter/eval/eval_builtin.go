package eval

import "ape/interpreter/data"

var builtins = map[string]*data.Builtin{
	"len": &data.Builtin{
		Fn: func(args ...data.Data) data.Data {
			if len(args) != 1 {
				return evalError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *data.String:
				return &data.Integer{Value: int64(len(arg.Value))}

			default:
				return evalError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
}

func evalBuiltin(value string) (*data.Builtin, bool) {
	result, ok := builtins[value]
	return result, ok
}
