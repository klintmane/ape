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

	case *ast.BlockStatement:
		return evalBlockStatement(node)

	case *ast.ReturnStatement:
		data := Eval(node.ReturnValue)
		if isError(data) {
			return data
		}
		return evalReturn(data)

	// Evaluate expressions
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}

		right := Eval(node.Right)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node)

		// Evaluate expressions
	case *ast.IntegerLiteral:
		return evalInteger(node.Value)

	case *ast.Boolean:
		return evalBoolean(node.Value)
	}

	return nil
}

func isError(d data.Data) bool {
	if d != nil {
		return d.Type() == data.ERROR_TYPE
	}
	return false
}
