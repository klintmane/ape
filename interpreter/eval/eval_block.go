package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalBlockStatement(block *ast.BlockStatement) data.Data {
	var result data.Data

	for _, statement := range block.Statements {
		result = Eval(statement)

		if result != nil {
			rt := result.Type()

			if rt == data.RETURN_TYPE || rt == data.ERROR_TYPE {
				return result
			}
		}
	}

	return result
}
