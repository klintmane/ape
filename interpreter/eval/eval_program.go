package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalProgram(program *ast.Program, env *data.Environment) data.Data {
	var result data.Data

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *data.Return:
			return result.Value

		case *data.Error:
			return result
		}
	}

	return result
}
