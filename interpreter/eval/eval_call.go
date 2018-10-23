package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalCallArguments(exps []ast.Expression, env *data.Environment) []data.Data {
	var result []data.Data

	for _, e := range exps {
		evaluated := Eval(e, env)

		if isError(evaluated) {
			return []data.Data{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalCallResult(function data.Data, args []data.Data) data.Data {
	fn, ok := function.(*data.Function)

	if !ok {
		return evalError("not a function: %s", fn.Type())
	}

	closure := evalCallClosure(fn, args)
	result := Eval(fn.Body, closure)

	return evalCallReturn(result)
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
