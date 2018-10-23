package eval

import "ape/interpreter/data"

func evalInfixExpression(
	operator string,
	left, right data.Data,
) data.Data {
	switch {
	case left.Type() == data.INTEGER_TYPE && right.Type() == data.INTEGER_TYPE:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return evalBoolean(left == right)
	case operator == "!=":
		return evalBoolean(left != right)
	default:
		return NULL
	}
}
