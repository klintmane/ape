package eval

import (
	"ape/src/ast"
	"ape/src/data"
)

func evalBlockStatement(block *ast.BlockStatement, env *data.Environment) data.Data {
	var result data.Data

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()

			if rt == data.RETURN_TYPE || rt == data.ERROR_TYPE {
				return result
			}
		}
	}

	return result
}
