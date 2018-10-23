package eval

import "ape/interpreter/data"

func evalInteger(value int64) *data.Integer {
	return &data.Integer{Value: value}
}
