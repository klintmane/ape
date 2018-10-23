package eval

import "ape/interpreter/data"

func evalIntegerInfixExpression(
	operator string,
	left, right data.Data,
) data.Data {
	leftVal := left.(*data.Integer).Value
	rightVal := right.(*data.Integer).Value

	switch operator {
	case "+":
		return &data.Integer{Value: leftVal + rightVal}
	case "-":
		return &data.Integer{Value: leftVal - rightVal}
	case "*":
		return &data.Integer{Value: leftVal * rightVal}
	case "/":
		return &data.Integer{Value: leftVal / rightVal}
	case "<":
		return evalBoolean(leftVal < rightVal)
	case ">":
		return evalBoolean(leftVal > rightVal)
	case "==":
		return evalBoolean(leftVal == rightVal)
	case "!=":
		return evalBoolean(leftVal != rightVal)
	default:
		return NULL
	}
}
