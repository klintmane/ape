package eval

import (
	"ape/interpreter/data"
	"fmt"
)

func _len(args ...data.Data) data.Data {
	if len(args) != 1 {
		return evalError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	switch arg := args[0].(type) {
	case *data.Array:
		return &data.Integer{Value: int64(len(arg.Elements))}

	case *data.String:
		return &data.Integer{Value: int64(len(arg.Value))}

	default:
		return evalError("argument to 'len' not supported, got %s", args[0].Type())
	}
}

func _head(args ...data.Data) data.Data {
	if len(args) != 1 {
		return evalError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	if args[0].Type() != data.ARRAY_TYPE {
		return evalError("argument to 'first' must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*data.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func _tail(args ...data.Data) data.Data {
	if len(args) != 1 {
		return evalError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	if args[0].Type() != data.ARRAY_TYPE {
		return evalError("argument to 'rest' must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*data.Array)
	length := len(arr.Elements)

	if length > 0 {
		newElements := make([]data.Data, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &data.Array{Elements: newElements}
	}

	return NULL
}

func _last(args ...data.Data) data.Data {
	if len(args) != 1 {
		return evalError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	if args[0].Type() != data.ARRAY_TYPE {
		return evalError("argument to 'last' must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*data.Array)
	length := len(arr.Elements)

	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func _push(args ...data.Data) data.Data {
	if len(args) != 2 {
		return evalError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	if args[0].Type() != data.ARRAY_TYPE {
		return evalError("argument to 'push' must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*data.Array)
	length := len(arr.Elements)
	newElements := make([]data.Data, length+1, length+1)

	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &data.Array{Elements: newElements}
}

func _print(args ...data.Data) data.Data {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return NULL
}

var builtins = map[string]*data.Builtin{
	"len": &data.Builtin{
		Fn: _len,
	},
	"head": &data.Builtin{
		Fn: _head,
	},
	"tail": &data.Builtin{
		Fn: _tail,
	},
	"last": &data.Builtin{
		Fn: _last,
	},
	"push": &data.Builtin{
		Fn: _push,
	},
	"print": &data.Builtin{
		Fn: _print,
	},
}

func evalBuiltin(value string) (*data.Builtin, bool) {
	result, ok := builtins[value]
	return result, ok
}
