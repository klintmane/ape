package eval

import (
	"ape/src/ast"
	"ape/src/data"
)

func Eval(node ast.Node, env *data.Environment) data.Data {
	switch node := node.(type) {

	// Evaluate statements
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.LetStatement:
		data := Eval(node.Value, env)
		if isError(data) {
			return data
		}
		env.Set(node.Name.Value, data)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.ReturnStatement:
		data := Eval(node.ReturnValue, env)
		if isError(data) {
			return data
		}
		return evalReturn(data)

	// Evaluate expressions
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

		// Evaluate expressions
	case *ast.IntegerLiteral:
		return evalInteger(node.Value)

	case *ast.StringLiteral:
		return evalString(node.Value)

	case *ast.Boolean:
		return evalBoolean(node.Value)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		return evalFunction(node, env)

	case *ast.CallExpression:
		fn := Eval(node.Function, env)
		if isError(fn) {
			return fn
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return evalCallResult(fn, args)

	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)

		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}

		return evalArray(elements)

	case *ast.HashLiteral:
		return evalHashLiteral(node, env)

	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}

		return evalIndexExpression(left, index)
	}

	return nil
}

func isError(d data.Data) bool {
	if d != nil {
		return d.Type() == data.ERROR_TYPE
	}
	return false
}
