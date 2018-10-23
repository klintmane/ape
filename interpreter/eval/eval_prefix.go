package eval

import (
	"ape/interpreter/data"
)

func evalPrefixExpression(operator string, right data.Data) data.Data {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}
