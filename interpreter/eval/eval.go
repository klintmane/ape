package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func Eval(node ast.Node) data.Data {
	switch node := node.(type) {

	// Evaluate statements
	case *ast.Program:
		return evalProgram(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

		// Evaluate expressions
	case *ast.IntegerLiteral:
		return evalInteger(node.Value)
	case *ast.Boolean:
		return evalBoolean(node.Value)
	}
	return nil
}
