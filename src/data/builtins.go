package data

import (
	"fmt"
)

var Builtins = []struct {
	Name       string
	Definition *Builtin
}{
	{"len", &Builtin{Fn: _len}},
	{"head", &Builtin{Fn: _head}},
	{"tail", &Builtin{Fn: _tail}},
	{"last", &Builtin{Fn: _last}},
	{"push", &Builtin{Fn: _push}},
	{"print", &Builtin{Fn: _print}},
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func GetBuiltinDef(name string) *Builtin {
	for _, b := range Builtins {
		if b.Name == name {
			return b.Definition
		}
	}
	return nil
}

func _len(args ...Data) Data {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Integer{Value: int64(len(arg.Elements))}
	case *String:
		return &Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to 'len' not supported, got %s", args[0].Type())
	}
}

func _head(args ...Data) Data {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != ARRAY_TYPE {
		return newError("argument to 'first' must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func _tail(args ...Data) Data {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != ARRAY_TYPE {
		return newError("argument to 'rest' must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)

	if length > 0 {
		newElements := make([]Data, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &Array{Elements: newElements}
	}

	return NULL
}

func _last(args ...Data) Data {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != ARRAY_TYPE {
		return newError("argument to 'last' must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)

	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func _push(args ...Data) Data {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].Type() != ARRAY_TYPE {
		return newError("argument to 'push' must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)
	newElements := make([]Data, length+1, length+1)

	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &Array{Elements: newElements}
}

func _print(args ...Data) Data {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return NULL
}
