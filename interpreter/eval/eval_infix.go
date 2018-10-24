package eval

import "ape/interpreter/data"

func evalInfixExpression(
	operator string,
	left, right data.Data,
) data.Data {
	switch {
	case left.Type() == data.INTEGER_TYPE && right.Type() == data.INTEGER_TYPE:
		return evalIntegerInfixExpression(operator, left, right)

	case left.Type() == data.STRING_TYPE && right.Type() == data.STRING_TYPE:
		return evalStringInfixExpression(operator, left, right)

	case left.Type() != right.Type():
		return evalError("Type mismatch: %s %s %s", left.Type(), operator, right.Type())

	case operator == "==":
		return evalBoolean(left == right)

	case operator == "!=":
		return evalBoolean(left != right)

	default:
		return evalError("Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
