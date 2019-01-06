package eval

import "github.com/ape-lang/ape/src/data"

func evalBangOperatorExpression(right data.Data) data.Data {
	switch right {
	case data.TRUE:
		return data.FALSE
	case data.FALSE:
		return data.TRUE
	case data.NULL:
		return data.TRUE
	default:
		return data.FALSE
	}
}
