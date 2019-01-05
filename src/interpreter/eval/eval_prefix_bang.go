package eval

import "github.com/ape-lang/ape/src/data"

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
