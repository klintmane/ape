package eval

import "ape/src/data"

func evalStringInfixExpression(operator string, left, right data.Data) data.Data {
	if operator != "+" {
		return evalError("Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}

	leftVal := left.(*data.String).Value
	rightVal := right.(*data.String).Value

	return &data.String{Value: leftVal + rightVal}
}
