package eval

import (
	"ape/src/interpreter/data"
)

func evalCallResult(function data.Data, args []data.Data) data.Data {
	switch fn := function.(type) {
	case *data.Function:
		closure := evalCallClosure(fn, args)
		result := Eval(fn.Body, closure)
		return evalCallReturn(result)

	case *data.Builtin:
		return fn.Fn(args...)

	default:
		return evalError("not a function: %s", fn.Type())
	}
}

func evalCallClosure(fn *data.Function, args []data.Data) *data.Environment {
	env := data.NewEnvironmentClosure(fn.Env)
	for i, param := range fn.Parameters {
		env.Set(param.Value, args[i])
	}
	return env
}

// Consume the Return (return the Return.Value instead of the Return) to avoid returning from the parent scope
func evalCallReturn(d data.Data) data.Data {
	if ret, ok := d.(*data.Return); ok {
		return ret.Value
	}
	return d
}
