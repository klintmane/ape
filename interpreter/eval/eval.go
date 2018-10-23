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
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

		// Evaluate expressions
	case *ast.IntegerLiteral:
		return evalInteger(node.Value)
	case *ast.Boolean:
		return evalBoolean(node.Value)
	}
	return nil
}
