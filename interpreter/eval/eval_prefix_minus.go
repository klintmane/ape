package eval

import "ape/interpreter/data"

func evalMinusPrefixOperatorExpression(right data.Data) data.Data {
	if right.Type() != data.INTEGER_TYPE {
		return NULL
	}

	value := right.(*data.Integer).Value
	return &data.Integer{Value: -value}
}
