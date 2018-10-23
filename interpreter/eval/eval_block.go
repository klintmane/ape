package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalBlockStatement(block *ast.BlockStatement) data.Data {
	var result data.Data
	for _, statement := range block.Statements {
		result = Eval(statement)
	}
	return result
}
