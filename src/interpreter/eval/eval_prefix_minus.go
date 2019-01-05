package eval

import "github.com/ape-lang/ape/src/data"

func evalMinusPrefixOperatorExpression(right data.Data) data.Data {
	if right.Type() != data.INTEGER_TYPE {
		return evalError("Unknown operator: -%s", right.Type())
	}

	value := right.(*data.Integer).Value
	return &data.Integer{Value: -value}
}
