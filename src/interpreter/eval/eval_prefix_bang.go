package eval

import "ape/src/interpreter/data"

func evalBangOperatorExpression(right data.Data) data.Data {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
