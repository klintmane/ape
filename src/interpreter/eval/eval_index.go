package eval

import "ape/src/data"

func evalIndexExpression(left, index data.Data) data.Data {
	switch {
	case left.Type() == data.ARRAY_TYPE && index.Type() == data.INTEGER_TYPE:
		return evalArrayIndexExpression(left, index)
	case left.Type() == data.HASH_TYPE:
		return evalHashIndexExpression(left, index)
	default:
		return evalError("index operator not supported: %s", left.Type())
	}
}
