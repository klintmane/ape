package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalProgram(program *ast.Program) data.Data {
	var result data.Data

	for _, statement := range program.Statements {
		result = Eval(statement)

		if ret, ok := result.(*data.Return); ok {
			return ret.Value
		}
	}

	return result
}
